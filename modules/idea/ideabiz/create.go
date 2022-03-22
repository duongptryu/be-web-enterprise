package ideabiz

import (
	"context"
	log "github.com/sirupsen/logrus"
	"time"
	"web/common"
	"web/components/mailprovider"
	"web/modules/acayear/acayearstore"
	"web/modules/category/categorymodel"
	"web/modules/category/categorystore"
	"web/modules/department/departmentstore"
	"web/modules/idea/ideamodel"
	"web/modules/idea/ideastore"
	"web/modules/notification/notificationmodel"
	"web/modules/notification/notificationstore"
	"web/modules/user/usermodel"
	"web/modules/user/userstore"
)

type createIdeaBiz struct {
	store             ideastore.IdeaStore
	categoryStore     categorystore.CategoryStore
	acaYearStore      acayearstore.AcademicYearStore
	userStore         userstore.UserStore
	departmentStore   departmentstore.DepartmentStore
	mailProvider      mailprovider.MailProvider
	notificationStore notificationstore.NotificationStore
}

func NewCreateIdeaBiz(store ideastore.IdeaStore, categoryStore categorystore.CategoryStore, acaYearStore acayearstore.AcademicYearStore, userStore userstore.UserStore, departmentStore departmentstore.DepartmentStore, mailProvider mailprovider.MailProvider, notificationStore notificationstore.NotificationStore) *createIdeaBiz {
	return &createIdeaBiz{
		store:             store,
		categoryStore:     categoryStore,
		acaYearStore:      acaYearStore,
		userStore:         userStore,
		departmentStore:   departmentStore,
		mailProvider:      mailProvider,
		notificationStore: notificationStore,
	}
}

func (biz *createIdeaBiz) CreateIdeaBiz(ctx context.Context, data *ideamodel.IdeaCreate) error {
	cateExist, err := biz.categoryStore.FindCategory(ctx, map[string]interface{}{"id": data.CategoryId})
	if err != nil {
		return err
	}
	if cateExist.Id == 0 {
		return common.ErrDataNotFound(categorymodel.EntityName)
	}

	acaExist, err := biz.acaYearStore.FindAcaYear(ctx, map[string]interface{}{"status": true})
	if err != nil {
		return err
	}
	if acaExist.Id == 0 {
		return ideamodel.ErrAcademicYearNotReady
	}

	timeNow := time.Now()
	if timeNow.After(acaExist.FirstClosureDate) {
		return ideamodel.ErrFirstClosureDateExpired
	}

	owner, err := biz.userStore.FindUser(ctx, map[string]interface{}{"id": data.UserId})
	if err != nil {
		return err
	}
	if owner.Id == 0 {
		return common.ErrDataNotFound(usermodel.EntityName)
	}

	if owner.Role != common.RoleStaff {
		return ideamodel.ErrAccountCannotCreateIdea
	}

	data.DepartmentId = owner.DepartmentId
	data.AcaYearId = acaExist.Id
	data.Status = true
	if err := biz.store.CreateIdea(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(ideamodel.EntityName, err)
	}

	//push noti email for QAC of this department
	go biz.pushNotiEmailForQAC(ctx, data, owner)

	return nil
}

func (biz *createIdeaBiz) pushNotiEmailForQAC(ctx context.Context, data *ideamodel.IdeaCreate, user *usermodel.User) {
	departmentDb, err := biz.departmentStore.FindDepartment(ctx, map[string]interface{}{"id": data.DepartmentId})
	if err != nil {
		log.Error(err)
		return
	}
	if departmentDb.Id == 0 {
		log.Error("Department not found with id - ", data.DepartmentId)
		return
	}

	ownerDepartment, err := biz.userStore.FindUser(ctx, map[string]interface{}{"id": departmentDb.LeaderId})
	if err != nil {
		log.Error(err)
		return
	}
	if ownerDepartment.Id == 0 {
		log.Error("Owner Department Not Found - ", departmentDb.LeaderId)
		return
	}

	go biz.mailProvider.SendMailNotifyNewIdea(ctx, &mailprovider.MailDataForIdea{
		Email:         ownerDepartment.Email,
		Name:          ownerDepartment.FullName,
		NameUserPush:  user.FullName,
		EmailUserPush: user.Email,
		Title:         data.Title,
		Id:            data.Id,
		Content:       data.Content,
		CreatedAt:     data.CreatedAt,
		ThumbnailUrl:  data.ThumbnailUrl,
	})

	newNoti := notificationmodel.NotificationIdeaCreate{
		TypeNoti: common.NewIdeaNotification,
		OwnerId:  ownerDepartment.Id,
		IdeaId:   data.Id,
		UserId:   data.UserId,
	}
	if err := biz.notificationStore.CreateNotification(ctx, &newNoti); err != nil {
		log.Error(err)
	}
}
