
��

blog.protoblog"

EmptyReq"
	EmptyResp"
IdReq
id (Rid"
IdsReq
ids (Rids"�
	PageQuery
page (Rpage
	page_size (RpageSize
sorts (	Rsorts

conditions (	R
conditions
args (	Rargs"<
	PageLimit
page (Rpage
	page_size (RpageSize"6
PageSort
field (	Rfield
order (	Rorder"m

PageCondition
field (	Rfield
value (	Rvalue
logic (	Rlogic
operator (	Roperator"Q
PageResp
page (Rpage
	page_size (RpageSize
total (Rtotal"0
	BatchResp#

success_count (RsuccessCount"!
	CountResp
count (Rcount"V
LoginReq
username (	Rusername
password (	Rpassword
code (	Rcode"�
	LoginResp
user_id (RuserId
username (	Rusername
nickname (	Rnickname
avatar (	Ravatar
intro (	Rintro
website (	Rwebsite
email (	Remail%
roles (2.blog.RoleLabelRroles"K
	RoleLabel
	role_name (	RroleName!
role_comment (	RroleComment"U

OauthLoginReq
platform (	Rplatform
code (	Rcode
state (	Rstate"%
OauthLoginUrlResp
url (	Rurl"^
ResetPasswordReq
username (	Rusername
password (	Rpassword
code (	Rcode"*
UserEmailReq
username (	Rusername"�
Api
id (Rid
name (	Rname
path (	Rpath
method (	Rmethod
	parent_id (RparentId
	traceable (R	traceable
status (Rstatus

created_at (R	createdAt

updated_at	 (R	updatedAt"�

ApiDetails
id (Rid
name (	Rname
path (	Rpath
method (	Rmethod
	parent_id (RparentId
	traceable (R	traceable
status (Rstatus

created_at (R	createdAt

updated_at	 (R	updatedAt,
children
 (2.blog.ApiDetailsRchildren"I
FindApiListResp
total (Rtotal$
list (2.blog.ApiDetailsRlist"�
Menu
id (Rid
	parent_id (RparentId
title (	Rtitle
path (	Rpath
name (	Rname
	component (	R	component
redirect (	Rredirect
type (Rtype
extra	 (	Rextra

created_at
 (R	createdAt

updated_at (R	updatedAt"�
MenuDetails
id (Rid
	parent_id (RparentId
title (	Rtitle
type (Rtype
path (	Rpath
name (	Rname
	component (	R	component
redirect (	Rredirect
extra	 (	Rextra

created_at
 (R	createdAt

updated_at (R	updatedAt-
children (2.blog.MenuDetailsRchildren"K
MenuPageResp
total (Rtotal%
list (2.blog.MenuDetailsRlist"6
SyncMenuReq'
menus (2.blog.MenuDetailsRmenus"�
Role
id (Rid
	parent_id (RparentId
role_domain (	R
roleDomain
	role_name (	RroleName!
role_comment (	RroleComment

is_disable (R	isDisable

is_default (R	isDefault

created_at (R	createdAt

updated_at	 (R	updatedAt"�
RoleDetails
id (Rid
	parent_id (RparentId
role_domain (	R
roleDomain
	role_name (	RroleName!
role_comment (	RroleComment

is_disable (R	isDisable

is_default (R	isDefault

created_at (R	createdAt

updated_at	 (R	updatedAt-
children
 (2.blog.RoleDetailsRchildren"K
RolePageResp
total (Rtotal%
list (2.blog.RoleDetailsRlist"`
RoleResourcesResp
role_id (RroleId
api_ids (RapiIds
menu_ids (RmenuIds"E
UpdateRoleApisReq
role_id (RroleId
api_ids (RapiIds"H
UpdateRoleMenusReq
role_id (RroleId
menu_ids (RmenuIds"�
LoginHistory
id (Rid

login_type (	R	loginType
agent (	Ragent

ip_address (	R	ipAddress
	ip_source (	RipSource

login_time (	R	loginTime"T
LoginHistoryPageResp
total (Rtotal&
list (2.blog.LoginHistoryRlist"�
UserInfoResp
id (Rid
user_id (RuserId
email (	Remail
nickname (	Rnickname
avatar (	Ravatar
phone (	Rphone
intro (	Rintro
website (	Rwebsite

created_at	 (R	createdAt

updated_at
 (R	updatedAt"�
User
id (Rid
username (	Rusername
email (	Remail
nickname (	Rnickname
avatar (	Ravatar
phone (	Rphone
intro (	Rintro
website (	Rwebsite
status	 (Rstatus#

register_type
 (	RregisterType

ip_address (	R	ipAddress
	ip_source (	RipSource

created_at
 (R	createdAt

updated_at (R	updatedAt%
roles (2.blog.RoleLabelRroles"H
UserInfoPageResp
total (Rtotal
list (2
.blog.UserRlist"�
UpdateUserInfoReq
user_id (RuserId
nickname (	Rnickname
phone (	Rphone
website (	Rwebsite
intro (	Rintro"F
UpdateUserAvatarReq
user_id (RuserId
avatar (	Ravatar"F
UpdateUserStatusReq
user_id (RuserId
status (Rstatus"G
UpdateUserRoleReq
user_id (RuserId
role_ids (RroleIds""
UserIdReq
user_id (RuserId".

FindConfigReq

config_key (	R	configKey"3
FindConfigResp!
config_value (	RconfigValue"Q

SaveConfigReq

config_key (	R	configKey!
config_value (	RconfigValue"�
Article
id (Rid
user_id (RuserId
category_id (R
categoryId#

article_cover (	RarticleCover#

article_title (	RarticleTitle'
article_content (	RarticleContent
type (Rtype!
original_url (	RoriginalUrl
is_top	 (RisTop
	is_delete
 (RisDelete
status (Rstatus

created_at (R	createdAt

updated_at
 (R	updatedAt"J
ArticlePageResp
total (Rtotal!
list (2
.blog.ArticleRlist".
FindArticleByTagReq
tag_ids (RtagIds"=
FindArticleByCategoryReq!
category_ids (RcategoryIds"}
Category
id (Rid#

category_name (	RcategoryName

created_at (R	createdAt

updated_at (R	updatedAt"L
CategoryPageResp
total (Rtotal"
list (2.blog.CategoryRlist"<
FindCategoryByNameReq#

category_name (	RcategoryName"n
Tag
id (Rid
tag_name (	RtagName

created_at (R	createdAt

updated_at (R	updatedAt"B
TagPageResp
total (Rtotal
list (2	.blog.TagRlist"/
FindTagArticleCountReq
tag_id (RtagId"-
FindTagByNameReq
tag_name (	RtagName"�

FriendLink
id (Rid
	link_name (	RlinkName
link_avatar (	R
linkAvatar!
link_address (	RlinkAddress

link_intro (	R	linkIntro

created_at (R	createdAt

updated_at (R	updatedAt"P
FriendLinkPageResp
total (Rtotal$
list (2.blog.FriendLinkRlist"�
Remark
id (Rid
nickname (	Rnickname
avatar (	Ravatar'
message_content (	RmessageContent

ip_address (	R	ipAddress
	ip_source (	RipSource
time (Rtime
	is_review (RisReview

created_at	 (R	createdAt

updated_at
 (R	updatedAt"H
RemarkPageResp
total (Rtotal
list (2.blog.RemarkRlist"�
Comment
id (Rid
topic_id (RtopicId
	parent_id (RparentId

session_id (R	sessionId
user_id (RuserId"

reply_user_id (RreplyUserId'
comment_content (	RcommentContent
type (Rtype
status	 (Rstatus
	is_review
 (RisReview

created_at (R	createdAt

updated_at (R	updatedAt"J
CommentPageResp
total (Rtotal!
list (2
.blog.CommentRlist"�
CommentReply
id (Rid
topic_id (RtopicId
	parent_id (RparentId

session_id (R	sessionId
user_id (RuserId"

reply_user_id (RreplyUserId'
comment_content (	RcommentContent
type (Rtype
status	 (Rstatus
	is_review
 (RisReview

created_at (R	createdAt

updated_at (R	updatedAt&
user
 (2.blog.UserInfoRespRuser1

reply_user (2.blog.UserInfoRespR	replyUser

like_count (R	likeCount"T
CommentReplyPageResp
total (Rtotal&
list (2.blog.CommentReplyRlist"�
Photo
id (Rid
album_id (RalbumId

photo_name (	R	photoName

photo_desc (	R	photoDesc
	photo_src (	RphotoSrc
	is_delete (RisDelete

created_at (R	createdAt

updated_at (R	updatedAt"F

PhotoPageResp
total (Rtotal
list (2.blog.PhotoRlist"�

PhotoAlbum
id (Rid

album_name (	R	albumName

album_desc (	R	albumDesc
album_cover (	R
albumCover
	is_delete (RisDelete
status (Rstatus

created_at (R	createdAt

updated_at (R	updatedAt"P
PhotoAlbumPageResp
total (Rtotal$
list (2.blog.PhotoAlbumRlist"�
Page
id (Rid
	page_name (	RpageName

page_label (	R	pageLabel

page_cover (	R	pageCover

created_at (R	createdAt

updated_at (R	updatedAt"D
PagePageResp
total (Rtotal
list (2
.blog.PageRlist"�
Talk
id (Rid
user_id (RuserId
content (	Rcontent
images (	Rimages
is_top (RisTop
status (Rstatus

created_at (R	createdAt

updated_at (R	updatedAt"�
TalkDetailsDTO
id (Rid
user_id (RuserId
nickname (	Rnickname
avatar (	Ravatar
content (	Rcontent
img_list (	RimgList
is_top (RisTop
status (Rstatus

like_count	 (R	likeCount#

comment_count
 (RcommentCount

created_at (R	createdAt

updated_at (R	updatedAt"D
TalkPageResp
total (Rtotal
list (2
.blog.TalkRlist"�
OperationLog
id (Rid
user_id (RuserId
nickname (	Rnickname

ip_address (	R	ipAddress
	ip_source (	RipSource

opt_module (	R	optModule
opt_desc (	RoptDesc
request_url (	R
requestUrl%
request_method	 (	R
requestMethod%
request_header
 (	R
requestHeader!
request_data (	RrequestData#

response_data (	RresponseData'
response_status
 (RresponseStatus
cost (	Rcost

created_at (R	createdAt

updated_at (R	updatedAt"T
OperationLogPageResp
total (Rtotal&
list (2.blog.OperationLogRlist"�

ChatRecord
id (Rid
user_id (RuserId
nickname (	Rnickname
avatar (	Ravatar
content (	Rcontent

ip_address (	R	ipAddress
	ip_source (	RipSource
type (Rtype

created_at	 (R	createdAt

updated_at
 (R	updatedAt"P
ChatRecordPageResp
total (Rtotal$
list (2.blog.ChatRecordRlist"�
UploadRecordReq
id (Rid
user_id (RuserId
label (	Rlabel
	file_name (	RfileName
	file_size (RfileSize
file_md5 (	RfileMd5
file_url (	RfileUrl"�
UploadRecordResp
id (Rid
user_id (RuserId
label (	Rlabel
	file_name (	RfileName
	file_size (RfileSize
file_md5 (	RfileMd5
file_url (	RfileUrl

created_at (R	createdAt

updated_at	 (R	updatedAt2�
AuthRpc(
Login.blog.LoginReq.blog.LoginResp)
Logout.blog.EmptyReq.blog.EmptyResp)
Logoff.blog.EmptyReq.blog.EmptyResp+
Register.blog.LoginReq.blog.EmptyResp4

RegisterEmail.blog.UserEmailReq.blog.EmptyResp:
ForgetPasswordEmail.blog.UserEmailReq.blog.EmptyResp8

ResetPassword.blog.ResetPasswordReq.blog.EmptyResp2

OauthLogin.blog.OauthLoginReq.blog.LoginRespD
GetOauthAuthorizeUrl.blog.OauthLoginReq.blog.OauthLoginUrlResp2�
ApiRpc!
	CreateApi	.blog.Api	.blog.Api!
	UpdateApi	.blog.Api	.blog.Api)
	DeleteApi.blog.IdReq.blog.BatchResp.

DeleteApiList.blog.IdsReq.blog.BatchResp!
FindApi.blog.IdReq	.blog.Api1
FindApiList.blog.PageQuery.blog.FindApiListResp.
SyncApiList.blog.EmptyReq.blog.BatchResp/
CleanApiList.blog.EmptyReq.blog.BatchResp2�
MenuRpc$

CreateMenu
.blog.Menu
.blog.Menu$

UpdateMenu
.blog.Menu
.blog.Menu*

DeleteMenu.blog.IdReq.blog.BatchResp/
DeleteMenuList.blog.IdsReq.blog.BatchResp#
FindMenu.blog.IdReq
.blog.Menu3
FindMenuList.blog.PageQuery.blog.MenuPageResp2
SyncMenuList.blog.SyncMenuReq.blog.BatchResp0

CleanMenuList.blog.EmptyReq.blog.BatchResp2�
RoleRpc$

CreateRole
.blog.Role
.blog.Role$

UpdateRole
.blog.Role
.blog.Role*

DeleteRole.blog.IdReq.blog.BatchResp/
DeleteRoleList.blog.IdsReq.blog.BatchResp#
FindRole.blog.IdReq
.blog.Role3
FindRoleList.blog.PageQuery.blog.RolePageResp9
FindRoleResources.blog.IdReq.blog.RoleResourcesResp<
UpdateRoleMenus.blog.UpdateRoleMenusReq.blog.EmptyResp:
UpdateRoleApis.blog.UpdateRoleApisReq.blog.EmptyResp2�
UserRpcG
FindUserLoginHistoryList.blog.PageQuery.blog.LoginHistoryPageResp;
DeleteUserLoginHistoryList.blog.IdsReq.blog.BatchResp/
GetUserApis
.blog.UserIdReq.blog.FindApiListResp1
GetUserMenus
.blog.UserIdReq.blog.MenuPageResp1
GetUserRoles
.blog.UserIdReq.blog.RolePageResp0
GetUserInfo
.blog.UserIdReq.blog.UserInfoResp=
UpdateUserInfo.blog.UpdateUserInfoReq.blog.UserInfoRespA
UpdateUserAvatar.blog.UpdateUserAvatarReq.blog.UserInfoResp>
UpdateUserStatus.blog.UpdateUserStatusReq.blog.EmptyResp:
UpdateUserRole.blog.UpdateUserRoleReq.blog.EmptyResp7
FindUserList.blog.PageQuery.blog.UserInfoPageResp2x
	ConfigRpc2

SaveConfig.blog.SaveConfigReq.blog.EmptyResp7

FindConfig.blog.FindConfigReq.blog.FindConfigResp2�

ArticleRpc-

CreateArticle
.blog.Article
.blog.Article-

UpdateArticle
.blog.Article
.blog.Article-

DeleteArticle.blog.IdReq.blog.BatchResp2
DeleteArticleList.blog.IdsReq.blog.BatchResp)
FindArticle.blog.IdReq
.blog.Article9
FindArticleList.blog.PageQuery.blog.ArticlePageResp4
FindArticleCount.blog.PageQuery.blog.CountRespD
FindArticleByTag.blog.FindArticleByTagReq.blog.ArticlePageRespN
FindArticleByCategory.blog.FindArticleByCategoryReq.blog.ArticlePageResp2�
CategoryRpc0
CreateCategory.blog.Category.blog.Category0
UpdateCategory.blog.Category.blog.Category.
DeleteCategory.blog.IdReq.blog.BatchResp3
DeleteCategoryList.blog.IdsReq.blog.BatchResp+
FindCategory.blog.IdReq.blog.Category;
FindCategoryList.blog.PageQuery.blog.CategoryPageResp5
FindCategoryCount.blog.PageQuery.blog.CountResp2�
TagRpc!
	CreateTag	.blog.Tag	.blog.Tag!
	UpdateTag	.blog.Tag	.blog.Tag)
	DeleteTag.blog.IdReq.blog.BatchResp.

DeleteTagList.blog.IdsReq.blog.BatchResp!
FindTag.blog.IdReq	.blog.Tag1
FindTagList.blog.PageQuery.blog.TagPageResp0
FindTagCount.blog.PageQuery.blog.CountRespD
FindTagArticleCount.blog.FindTagArticleCountReq.blog.CountResp8
FindTagListByArticleId.blog.IdReq.blog.TagPageResp2�

FriendLinkRpc6
CreateFriendLink.blog.FriendLink.blog.FriendLink6
UpdateFriendLink.blog.FriendLink.blog.FriendLink0
DeleteFriendLink.blog.IdReq.blog.BatchResp5
DeleteFriendLinkList.blog.IdsReq.blog.BatchResp/
FindFriendLink.blog.IdReq.blog.FriendLink?
FindFriendLinkList.blog.PageQuery.blog.FriendLinkPageResp7
FindFriendLinkCount.blog.PageQuery.blog.CountResp2�
	remarkRpc*
CreateRemark.blog.Remark.blog.Remark*
UpdateRemark.blog.Remark.blog.Remark,
DeleteRemark.blog.IdReq.blog.BatchResp1
DeleteRemarkList.blog.IdsReq.blog.BatchResp'

FindRemark.blog.IdReq.blog.Remark7
FindRemarkList.blog.PageQuery.blog.RemarkPageResp3
FindRemarkCount.blog.PageQuery.blog.CountResp2�

commentRpc-

CreateComment
.blog.Comment
.blog.Comment-

UpdateComment
.blog.Comment
.blog.Comment-

DeleteComment.blog.IdReq.blog.BatchResp2
DeleteCommentList.blog.IdsReq.blog.BatchResp)
FindComment.blog.IdReq
.blog.Comment9
FindCommentList.blog.PageQuery.blog.CommentPageRespC
FindCommentReplyList.blog.PageQuery.blog.CommentReplyPageResp4
FindCommentCount.blog.PageQuery.blog.CountResp+
LikeComment.blog.IdReq.blog.EmptyResp2�
photoRpc'
CreatePhoto.blog.Photo.blog.Photo'
UpdatePhoto.blog.Photo.blog.Photo+
DeletePhoto.blog.IdReq.blog.BatchResp0
DeletePhotoList.blog.IdsReq.blog.BatchResp%
	FindPhoto.blog.IdReq.blog.Photo5

FindPhotoList.blog.PageQuery.blog.PhotoPageResp2
FindPhotoCount.blog.PageQuery.blog.CountResp6
CreatePhotoAlbum.blog.PhotoAlbum.blog.PhotoAlbum6
UpdatePhotoAlbum.blog.PhotoAlbum.blog.PhotoAlbum0
DeletePhotoAlbum.blog.IdReq.blog.BatchResp5
DeletePhotoAlbumList.blog.IdsReq.blog.BatchResp/
FindPhotoAlbum.blog.IdReq.blog.PhotoAlbum?
FindPhotoAlbumList.blog.PageQuery.blog.PhotoAlbumPageResp7
FindPhotoAlbumCount.blog.PageQuery.blog.CountResp2�
pageRpc$

CreatePage
.blog.Page
.blog.Page$

UpdatePage
.blog.Page
.blog.Page*

DeletePage.blog.IdReq.blog.BatchResp/
DeletePageList.blog.IdsReq.blog.BatchResp#
FindPage.blog.IdReq
.blog.Page3
FindPageList.blog.PageQuery.blog.PagePageResp1

FindPageCount.blog.PageQuery.blog.CountResp2�
talkRpc$

CreateTalk
.blog.Talk
.blog.Talk$

UpdateTalk
.blog.Talk
.blog.Talk*

DeleteTalk.blog.IdReq.blog.BatchResp/
DeleteTalkList.blog.IdsReq.blog.BatchResp#
FindTalk.blog.IdReq
.blog.Talk3
FindTalkList.blog.PageQuery.blog.TalkPageResp1

FindTalkCount.blog.PageQuery.blog.CountResp(
LikeTalk.blog.IdReq.blog.EmptyResp2�
logRpc<
CreateOperationLog.blog.OperationLog.blog.OperationLog<
UpdateOperationLog.blog.OperationLog.blog.OperationLog2
DeleteOperationLog.blog.IdReq.blog.BatchResp7
DeleteOperationLogList.blog.IdsReq.blog.BatchResp3
FindOperationLog.blog.IdReq.blog.OperationLogC
FindOperationLogList.blog.PageQuery.blog.OperationLogPageResp9
FindOperationLogCount.blog.PageQuery.blog.CountResp2�
chatRpc6
CreateChatRecord.blog.ChatRecord.blog.ChatRecord6
UpdateChatRecord.blog.ChatRecord.blog.ChatRecord0
DeleteChatRecord.blog.IdReq.blog.BatchResp5
DeleteChatRecordList.blog.IdsReq.blog.BatchResp/
FindChatRecord.blog.IdReq.blog.ChatRecord?
FindChatRecordList.blog.PageQuery.blog.ChatRecordPageResp7
FindChatRecordCount.blog.PageQuery.blog.CountResp2�
	uploadRpc;

UploadFile.blog.UploadRecordReq.blog.UploadRecordResp<
UploadVoice.blog.UploadRecordReq.blog.UploadRecordRespBZ./blogbproto3
