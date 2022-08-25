package proto

import (
	"encoding/gob"
)

// 注册需要转发到GS的消息
func init() {
	gob.Register(&NullMsg{})
	gob.Register(&PlayerSetPauseReq{})
	gob.Register(&SetPlayerBornDataReq{})
	gob.Register(&PlayerDataNotify{})
	gob.Register(&StoreWeightLimitNotify{})
	gob.Register(&PlayerStoreNotify{})
	gob.Register(&Item_Equip{})
	gob.Register(&Equip_Weapon{})
	gob.Register(&AvatarDataNotify{})
	gob.Register(&PlayerEnterSceneNotify{})
	gob.Register(&OpenStateUpdateNotify{})
	gob.Register(&PropValue_Ival{})
	gob.Register(&GetPlayerSocialDetailReq{})
	gob.Register(&GetPlayerSocialDetailRsp{})
	gob.Register(&EnterSceneReadyReq{})
	gob.Register(&EnterScenePeerNotify{})
	gob.Register(&EnterSceneReadyRsp{})
	gob.Register(&GetScenePointReq{})
	gob.Register(&GetScenePointRsp{})
	gob.Register(&GetSceneAreaReq{})
	gob.Register(&GetSceneAreaRsp{})
	gob.Register(&ServerTimeNotify{})
	gob.Register(&WorldPlayerInfoNotify{})
	gob.Register(&WorldDataNotify{})
	gob.Register(&PlayerWorldSceneInfoListNotify{})
	gob.Register(&HostPlayerNotify{})
	gob.Register(&SceneTimeNotify{})
	gob.Register(&PlayerGameTimeNotify{})
	gob.Register(&PlayerEnterSceneInfoNotify{})
	gob.Register(&SceneAreaWeatherNotify{})
	gob.Register(&ScenePlayerInfoNotify{})
	gob.Register(&SceneTeamUpdateNotify{})
	gob.Register(&SceneEntityInfo_Avatar{})
	gob.Register(&SyncTeamEntityNotify{})
	gob.Register(&SyncScenePlayTeamEntityNotify{})
	gob.Register(&SceneInitFinishRsp{})
	gob.Register(&EnterSceneDoneRsp{})
	gob.Register(&PlayerTimeNotify{})
	gob.Register(&SceneEntityAppearNotify{})
	gob.Register(&WorldPlayerLocationNotify{})
	gob.Register(&ScenePlayerLocationNotify{})
	gob.Register(&WorldPlayerRTTNotify{})
	gob.Register(&EnterWorldAreaReq{})
	gob.Register(&EnterWorldAreaRsp{})
	gob.Register(&PostEnterSceneRsp{})
	gob.Register(&TowerAllDataRsp{})
	gob.Register(&SceneTransToPointReq{})
	gob.Register(&SceneTransToPointRsp{})
	gob.Register(&SceneEntityDisappearNotify{})
	gob.Register(&CombatInvocationsNotify{})
	gob.Register(&MarkMapReq{})
	gob.Register(&ChangeAvatarReq{})
	gob.Register(&ChangeAvatarRsp{})
	gob.Register(&SetUpAvatarTeamReq{})
	gob.Register(&SetUpAvatarTeamRsp{})
	gob.Register(&AvatarTeamUpdateNotify{})
	gob.Register(&ChooseCurAvatarTeamReq{})
	gob.Register(&ChooseCurAvatarTeamRsp{})
	gob.Register(&StoreItemChangeNotify{})
	gob.Register(&ItemAddHintNotify{})
	gob.Register(&StoreItemDelNotify{})
	gob.Register(&PlayerPropNotify{})
	gob.Register(&Item_Material{})
	gob.Register(&GetGachaInfoRsp{})
	gob.Register(&DoGachaReq{})
	gob.Register(&DoGachaRsp{})
	gob.Register(&Item_Furniture{})
	gob.Register(&SceneEntityInfo_Monster{})
	gob.Register(&EntityFightPropUpdateNotify{})
	gob.Register(&QueryPathReq{})
	gob.Register(&QueryPathRsp{})
	gob.Register(&PingReq{})
	gob.Register(&PingRsp{})
	gob.Register(&AbilityInvocationsNotify{})
	gob.Register(&ClientAbilityInitFinishNotify{})
	gob.Register(&EntityAiSyncNotify{})
	gob.Register(&AvatarFightPropNotify{})
	gob.Register(&AvatarEquipChangeNotify{})
	gob.Register(&AvatarAddNotify{})
	gob.Register(&WearEquipReq{})
	gob.Register(&WearEquipRsp{})
	gob.Register(&ChangeGameTimeReq{})
	gob.Register(&ChangeGameTimeRsp{})
	gob.Register(&SetPlayerBirthdayReq{})
	gob.Register(&SetPlayerBirthdayRsp{})
	gob.Register(&SetNameCardReq{})
	gob.Register(&SetNameCardRsp{})
	gob.Register(&SetPlayerSignatureReq{})
	gob.Register(&SetPlayerSignatureRsp{})
	gob.Register(&SetPlayerNameReq{})
	gob.Register(&SetPlayerNameRsp{})
	gob.Register(&SetPlayerHeadImageReq{})
	gob.Register(&SetPlayerHeadImageRsp{})
	gob.Register(&GetAllUnlockNameCardRsp{})
	gob.Register(&UnlockNameCardNotify{})
	gob.Register(&AskAddFriendReq{})
	gob.Register(&DealAddFriendReq{})
	gob.Register(&GetPlayerFriendListRsp{})
	gob.Register(&GetPlayerAskFriendListRsp{})
	gob.Register(&AskAddFriendRsp{})
	gob.Register(&AskAddFriendNotify{})
	gob.Register(&DealAddFriendRsp{})
	gob.Register(&GetOnlinePlayerListReq{})
	gob.Register(&GetOnlinePlayerListRsp{})
	gob.Register(&DelTeamEntityNotify{})
	gob.Register(&PlayerApplyEnterMpReq{})
	gob.Register(&PlayerApplyEnterMpRsp{})
	gob.Register(&PlayerApplyEnterMpNotify{})
	gob.Register(&PlayerApplyEnterMpResultReq{})
	gob.Register(&PlayerApplyEnterMpResultRsp{})
	gob.Register(&PlayerApplyEnterMpResultNotify{})
	gob.Register(&PlayerGetForceQuitBanInfoRsp{})
}
