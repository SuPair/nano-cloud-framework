package framework
// FFFF Operate: FF resource: FF type


const (
	ResourceOffset = 8
	OperateOffset = 8 + ResourceOffset
)
//Offset
const (
	MessageRequest = iota
	MessageResponse
	MessageEvent
)

type ServiceType uint
//Service Type
const (
	ServiceTypeCore = iota
	ServiceTypeCell
	ServiceTypeImage
)

const (
	ServiceTypeStringCore = "core"
)

//resource
const (
	ResourceService = iota
	ResourceConnection
	ResourceComputePool
	ResourceComputeCell
	ResourceHost
	ResourceGuest
	ResourceInstance
	ResourceZone
	ResourceMediaImage
	ResourceDiskImage
	ResourceImageServer
	ResourceAddress
	ResourceCore
	ResourceMemory
	ResourceDisk
	ResourceAuth
	ResourceSnapshot
	ResourceMedia
	ResourceStoragePool
	ResourceMigration
	ResourceAddressPool
	ResourceAddressRange
	ResourceSystem
	ResourceName
	ResourcePriority
	ResourceDiskThreshold
	ResourceNetworkThreshold
)


//resource operate
const (
	OperateCreate = iota
	OperateDelete
	OperateModify
	OperateQuery
	OperateAdd
	OperateRemove
	OperateRegister
	OperateUnregiser
	OperateQueryDetail
	OperateQueryUnallocated
	OperateGet
	OperateQueryStatus
	OperateStart
	OperateStop
	OperateGetStatus
	OperateResize
	OperateShrink
	OperateAttach
	OperateDetach
	OperateCommit
	OperatePull
	OperateRestore
	OperateEnable
	OperateDisable
	OperateMigrate
	OperatePurge
	OperateReset
	OperateStartBatchCreate
	OperateGetBatchCreate
	OperateStartBatchDelete
	OperateGetBatchDelete
	OperateStartBatchStop
	OperateGetBatchStop
)


const (
	RegisterServiceRequest   = OperateRegister<<OperateOffset | ResourceService<<ResourceOffset | MessageRequest
	RegisterServiceResponse  = OperateRegister<<OperateOffset | ResourceService<<ResourceOffset | MessageResponse
	CreateConnectionRequest  = OperateCreate<<OperateOffset | ResourceConnection<<ResourceOffset | MessageRequest
	CreateConnectionResponse = OperateCreate<<OperateOffset | ResourceConnection<<ResourceOffset | MessageResponse
	DeleteConnectionRequest  = OperateDelete<<OperateOffset | ResourceConnection<<ResourceOffset | MessageRequest
	DeleteConnectionResponse = OperateDelete<<OperateOffset | ResourceConnection<<ResourceOffset | MessageResponse

	//zone
	QueryZoneRequest        = OperateQuery<<OperateOffset | ResourceZone<<ResourceOffset | MessageRequest
	QueryZoneResponse       = OperateQuery<<OperateOffset | ResourceZone<<ResourceOffset | MessageResponse
	QueryZoneStatusRequest  = OperateQueryStatus<<OperateOffset | ResourceZone<<ResourceOffset | MessageRequest
	QueryZoneStatusResponse = OperateQueryStatus<<OperateOffset | ResourceZone<<ResourceOffset | MessageResponse

	//compute pool
	QueryComputePoolRequest  = OperateQuery<<OperateOffset | ResourceComputePool<<ResourceOffset | MessageRequest
	QueryComputePoolResponse = OperateQuery<<OperateOffset | ResourceComputePool<<ResourceOffset | MessageResponse
	GetComputePoolRequest    = OperateGet<<OperateOffset | ResourceComputePool<<ResourceOffset | MessageRequest
	GetComputePoolResponse   = OperateGet<<OperateOffset | ResourceComputePool<<ResourceOffset | MessageResponse

	QueryComputePoolStatusRequest  = OperateQueryStatus<<OperateOffset | ResourceComputePool<<ResourceOffset | MessageRequest
	QueryComputePoolStatusResponse = OperateQueryStatus<<OperateOffset | ResourceComputePool<<ResourceOffset | MessageResponse
	GetComputePoolStatusRequest    = OperateGetStatus<<OperateOffset | ResourceComputePool<<ResourceOffset | MessageRequest
	GetComputePoolStatusResponse   = OperateGetStatus<<OperateOffset | ResourceComputePool<<ResourceOffset | MessageResponse

	QueryComputePoolDetailRequest  = OperateQueryDetail<<OperateOffset | ResourceComputePool<<ResourceOffset | MessageRequest
	QueryComputePoolDetailResponse = OperateQueryDetail<<OperateOffset | ResourceComputePool<<ResourceOffset | MessageResponse

	CreateComputePoolRequest  = OperateCreate<<OperateOffset | ResourceComputePool<<ResourceOffset | MessageRequest
	CreateComputePoolResponse = OperateCreate<<OperateOffset | ResourceComputePool<<ResourceOffset | MessageResponse
	DeleteComputePoolRequest  = OperateDelete<<OperateOffset | ResourceComputePool<<ResourceOffset | MessageRequest
	DeleteComputePoolResponse = OperateDelete<<OperateOffset | ResourceComputePool<<ResourceOffset | MessageResponse
	ModifyComputePoolRequest  = OperateModify<<OperateOffset | ResourceComputePool<<ResourceOffset | MessageRequest
	ModifyComputePoolResponse = OperateModify<<OperateOffset | ResourceComputePool<<ResourceOffset | MessageResponse

	//storage pool
	QueryStoragePoolRequest  = OperateQuery<<OperateOffset | ResourceStoragePool<<ResourceOffset | MessageRequest
	QueryStoragePoolResponse = OperateQuery<<OperateOffset | ResourceStoragePool<<ResourceOffset | MessageResponse
	GetStoragePoolRequest    = OperateGet<<OperateOffset | ResourceStoragePool<<ResourceOffset | MessageRequest
	GetStoragePoolResponse   = OperateGet<<OperateOffset | ResourceStoragePool<<ResourceOffset | MessageResponse

	QueryStoragePoolStatusRequest  = OperateQueryStatus<<OperateOffset | ResourceStoragePool<<ResourceOffset | MessageRequest
	QueryStoragePoolStatusResponse = OperateQueryStatus<<OperateOffset | ResourceStoragePool<<ResourceOffset | MessageResponse
	GetStoragePoolStatusRequest    = OperateGetStatus<<OperateOffset | ResourceStoragePool<<ResourceOffset | MessageRequest
	GetStoragePoolStatusResponse   = OperateGetStatus<<OperateOffset | ResourceStoragePool<<ResourceOffset | MessageResponse

	CreateStoragePoolRequest  = OperateCreate<<OperateOffset | ResourceStoragePool<<ResourceOffset | MessageRequest
	CreateStoragePoolResponse = OperateCreate<<OperateOffset | ResourceStoragePool<<ResourceOffset | MessageResponse
	DeleteStoragePoolRequest  = OperateDelete<<OperateOffset | ResourceStoragePool<<ResourceOffset | MessageRequest
	DeleteStoragePoolResponse = OperateDelete<<OperateOffset | ResourceStoragePool<<ResourceOffset | MessageResponse
	ModifyStoragePoolRequest  = OperateModify<<OperateOffset | ResourceStoragePool<<ResourceOffset | MessageRequest
	ModifyStoragePoolResponse = OperateModify<<OperateOffset | ResourceStoragePool<<ResourceOffset | MessageResponse

	//compute pool cell
	QueryComputePoolCellRequest  = OperateQuery<<OperateOffset | ResourceComputeCell<<ResourceOffset | MessageRequest
	QueryComputePoolCellResponse = OperateQuery<<OperateOffset | ResourceComputeCell<<ResourceOffset | MessageResponse
	GetComputePoolCellRequest    = OperateGet<<OperateOffset | ResourceComputeCell<<ResourceOffset | MessageRequest
	GetComputePoolCellResponse   = OperateGet<<OperateOffset | ResourceComputeCell<<ResourceOffset | MessageResponse

	QueryComputePoolCellStatusRequest  = OperateQueryStatus<<OperateOffset | ResourceComputeCell<<ResourceOffset | MessageRequest
	QueryComputePoolCellStatusResponse = OperateQueryStatus<<OperateOffset | ResourceComputeCell<<ResourceOffset | MessageResponse
	GetComputePoolCellStatusRequest    = OperateGetStatus<<OperateOffset | ResourceComputeCell<<ResourceOffset | MessageRequest
	GetComputePoolCellStatusResponse   = OperateGetStatus<<OperateOffset | ResourceComputeCell<<ResourceOffset | MessageResponse

	QueryComputePoolCellDetailRequest       = OperateQueryDetail<<OperateOffset | ResourceComputeCell<<ResourceOffset | MessageRequest
	QueryComputePoolCellDetailResponse      = OperateQueryDetail<<OperateOffset | ResourceComputeCell<<ResourceOffset | MessageResponse
	QueryUnallocatedComputePoolCellRequest  = OperateQueryUnallocated<<OperateOffset | ResourceComputeCell<<ResourceOffset | MessageRequest
	QueryUnallocatedComputePoolCellResponse = OperateQueryUnallocated<<OperateOffset | ResourceComputeCell<<ResourceOffset | MessageResponse
	AddComputePoolCellRequest               = OperateAdd<<OperateOffset | ResourceComputeCell<<ResourceOffset | MessageRequest
	AddComputePoolCellResponse              = OperateAdd<<OperateOffset | ResourceComputeCell<<ResourceOffset | MessageResponse
	RemoveComputePoolCellRequest            = OperateRemove<<OperateOffset | ResourceComputeCell<<ResourceOffset | MessageRequest
	RemoveComputePoolCellResponse           = OperateRemove<<OperateOffset | ResourceComputeCell<<ResourceOffset | MessageResponse
	ModifyComputePoolCellRequest            = OperateModify<<OperateOffset | ResourceComputeCell<<ResourceOffset | MessageRequest
	ModifyComputePoolCellResponse           = OperateModify<<OperateOffset | ResourceComputeCell<<ResourceOffset | MessageResponse
	EnableComputePoolCellRequest            = OperateEnable<<OperateOffset | ResourceComputeCell<<ResourceOffset | MessageRequest
	EnableComputePoolCellResponse           = OperateEnable<<OperateOffset | ResourceComputeCell<<ResourceOffset | MessageResponse
	DisableComputePoolCellRequest           = OperateDisable<<OperateOffset | ResourceComputeCell<<ResourceOffset | MessageRequest
	DisableComputePoolCellResponse          = OperateDisable<<OperateOffset | ResourceComputeCell<<ResourceOffset | MessageResponse

	//guest
	QueryGuestRequest  = OperateQuery<<OperateOffset | ResourceGuest<<ResourceOffset | MessageRequest
	QueryGuestResponse = OperateQuery<<OperateOffset | ResourceGuest<<ResourceOffset | MessageResponse

	GetGuestRequest  = OperateGet<<OperateOffset | ResourceGuest<<ResourceOffset | MessageRequest
	GetGuestResponse = OperateGet<<OperateOffset | ResourceGuest<<ResourceOffset | MessageResponse

	CreateGuestRequest  = OperateCreate<<OperateOffset | ResourceGuest<<ResourceOffset | MessageRequest
	CreateGuestResponse = OperateCreate<<OperateOffset | ResourceGuest<<ResourceOffset | MessageResponse

	DeleteGuestRequest  = OperateDelete<<OperateOffset | ResourceGuest<<ResourceOffset | MessageRequest
	DeleteGuestResponse = OperateDelete<<OperateOffset | ResourceGuest<<ResourceOffset | MessageResponse

	ModifyGuestRequest  = OperateModify<<OperateOffset | ResourceGuest<<ResourceOffset | MessageRequest
	ModifyGuestResponse = OperateModify<<OperateOffset | ResourceGuest<<ResourceOffset | MessageResponse

	ResetSystemRequest  = OperateReset<<OperateOffset | ResourceSystem<<ResourceOffset | MessageRequest
	ResetSystemResponse = OperateReset<<OperateOffset | ResourceSystem<<ResourceOffset | MessageResponse

	//instance

	QueryInstanceStatusRequest  = OperateQueryStatus<<OperateOffset | ResourceInstance<<ResourceOffset | MessageRequest
	QueryInstanceStatusResponse = OperateQueryStatus<<OperateOffset | ResourceInstance<<ResourceOffset | MessageResponse
	GetInstanceStatusRequest    = OperateGetStatus<<OperateOffset | ResourceInstance<<ResourceOffset | MessageRequest
	GetInstanceStatusResponse   = OperateGetStatus<<OperateOffset | ResourceInstance<<ResourceOffset | MessageResponse

	StartInstanceRequest  = OperateStart<<OperateOffset | ResourceInstance<<ResourceOffset | MessageRequest
	StartInstanceResponse = OperateStart<<OperateOffset | ResourceInstance<<ResourceOffset | MessageResponse

	StopInstanceRequest  = OperateStop<<OperateOffset | ResourceInstance<<ResourceOffset | MessageRequest
	StopInstanceResponse = OperateStop<<OperateOffset | ResourceInstance<<ResourceOffset | MessageResponse

	ModifyInstanceRequest  = OperateModify<<OperateOffset | ResourceInstance<<ResourceOffset | MessageRequest
	ModifyInstanceResponse = OperateModify<<OperateOffset | ResourceInstance<<ResourceOffset | MessageResponse

	MigrateInstanceRequest  = OperateMigrate<<OperateOffset | ResourceInstance<<ResourceOffset | MessageRequest
	MigrateInstanceResponse = OperateMigrate<<OperateOffset | ResourceInstance<<ResourceOffset | MessageResponse
	AttachInstanceRequest  = OperateAttach<<OperateOffset | ResourceInstance<<ResourceOffset | MessageRequest
	AttachInstanceResponse = OperateAttach<<OperateOffset | ResourceInstance<<ResourceOffset | MessageResponse

	DetachInstanceRequest  = OperateDetach<<OperateOffset | ResourceInstance<<ResourceOffset | MessageRequest
	DetachInstanceResponse = OperateDetach<<OperateOffset | ResourceInstance<<ResourceOffset | MessageResponse
	PurgeInstanceRequest  = OperatePurge<<OperateOffset | ResourceInstance<<ResourceOffset | MessageRequest
	PurgeInstanceResponse = OperatePurge<<OperateOffset | ResourceInstance<<ResourceOffset | MessageResponse

	//Media Image

	QueryMediaImageRequest  = OperateQuery<<OperateOffset | ResourceMediaImage<<ResourceOffset | MessageRequest
	QueryMediaImageResponse = OperateQuery<<OperateOffset | ResourceMediaImage<<ResourceOffset | MessageResponse

	GetMediaImageRequest  = OperateGet<<OperateOffset | ResourceMediaImage<<ResourceOffset | MessageRequest
	GetMediaImageResponse = OperateGet<<OperateOffset | ResourceMediaImage<<ResourceOffset | MessageResponse

	CreateMediaImageRequest  = OperateCreate<<OperateOffset | ResourceMediaImage<<ResourceOffset | MessageRequest
	CreateMediaImageResponse = OperateCreate<<OperateOffset | ResourceMediaImage<<ResourceOffset | MessageResponse

	DeleteMediaImageRequest  = OperateDelete<<OperateOffset | ResourceMediaImage<<ResourceOffset | MessageRequest
	DeleteMediaImageResponse = OperateDelete<<OperateOffset | ResourceMediaImage<<ResourceOffset | MessageResponse

	ModifyMediaImageRequest  = OperateModify<<OperateOffset | ResourceMediaImage<<ResourceOffset | MessageRequest
	ModifyMediaImageResponse = OperateModify<<OperateOffset | ResourceMediaImage<<ResourceOffset | MessageResponse

	//Disk Image

	QueryDiskImageRequest  = OperateQuery<<OperateOffset | ResourceDiskImage<<ResourceOffset | MessageRequest
	QueryDiskImageResponse = OperateQuery<<OperateOffset | ResourceDiskImage<<ResourceOffset | MessageResponse

	GetDiskImageRequest  = OperateGet<<OperateOffset | ResourceDiskImage<<ResourceOffset | MessageRequest
	GetDiskImageResponse = OperateGet<<OperateOffset | ResourceDiskImage<<ResourceOffset | MessageResponse

	CreateDiskImageRequest  = OperateCreate<<OperateOffset | ResourceDiskImage<<ResourceOffset | MessageRequest
	CreateDiskImageResponse = OperateCreate<<OperateOffset | ResourceDiskImage<<ResourceOffset | MessageResponse

	DeleteDiskImageRequest  = OperateDelete<<OperateOffset | ResourceDiskImage<<ResourceOffset | MessageRequest
	DeleteDiskImageResponse = OperateDelete<<OperateOffset | ResourceDiskImage<<ResourceOffset | MessageResponse

	ModifyDiskImageRequest  = OperateModify<<OperateOffset | ResourceDiskImage<<ResourceOffset | MessageRequest
	ModifyDiskImageResponse = OperateModify<<OperateOffset | ResourceDiskImage<<ResourceOffset | MessageResponse

	ModifyGuestNameRequest  = OperateModify<<OperateOffset | ResourceName<<ResourceOffset | MessageRequest
	ModifyGuestNameResponse = OperateModify<<OperateOffset | ResourceName<<ResourceOffset | MessageResponse

	ModifyCoreRequest  = OperateModify<<OperateOffset | ResourceCore<<ResourceOffset | MessageRequest
	ModifyCoreResponse = OperateModify<<OperateOffset | ResourceCore<<ResourceOffset | MessageResponse

	ModifyMemoryRequest  = OperateModify<<OperateOffset | ResourceMemory<<ResourceOffset | MessageRequest
	ModifyMemoryResponse = OperateModify<<OperateOffset | ResourceMemory<<ResourceOffset | MessageResponse

	ModifyPriorityRequest  = OperateModify << OperateOffset | ResourcePriority << ResourceOffset | MessageRequest
	ModifyPriorityResponse = OperateModify << OperateOffset | ResourcePriority << ResourceOffset | MessageResponse

	ModifyDiskThresholdRequest  = OperateModify << OperateOffset | ResourceDiskThreshold << ResourceOffset | MessageRequest
	ModifyDiskThresholdResponse = OperateModify << OperateOffset | ResourceDiskThreshold << ResourceOffset | MessageResponse

	ModifyNetworkThresholdRequest  = OperateModify << OperateOffset | ResourceNetworkThreshold << ResourceOffset | MessageRequest
	ModifyNetworkThresholdResponse = OperateModify << OperateOffset | ResourceNetworkThreshold << ResourceOffset | MessageResponse

	ModifyAuthRequest  = OperateModify<<OperateOffset | ResourceAuth<<ResourceOffset | MessageRequest
	ModifyAuthResponse = OperateModify<<OperateOffset | ResourceAuth<<ResourceOffset | MessageResponse

	GetAuthRequest  = OperateGet<<OperateOffset | ResourceAuth<<ResourceOffset | MessageRequest
	GetAuthResponse = OperateGet<<OperateOffset | ResourceAuth<<ResourceOffset | MessageResponse

	ResizeDiskRequest  = OperateResize<<OperateOffset | ResourceDisk<<ResourceOffset | MessageRequest
	ResizeDiskResponse = OperateResize<<OperateOffset | ResourceDisk<<ResourceOffset | MessageResponse

	ShrinkDiskRequest  = OperateShrink<<OperateOffset | ResourceDisk<<ResourceOffset | MessageRequest
	ShrinkDiskResponse = OperateShrink<<OperateOffset | ResourceDisk<<ResourceOffset | MessageResponse

	//instance media
	InsertMediaRequest  = OperateAttach<<OperateOffset | ResourceMedia<<ResourceOffset | MessageRequest
	InsertMediaResponse = OperateAttach<<OperateOffset | ResourceMedia<<ResourceOffset | MessageResponse
	EjectMediaRequest   = OperateDetach<<OperateOffset | ResourceMedia<<ResourceOffset | MessageRequest
	EjectMediaResponse  = OperateDetach<<OperateOffset | ResourceMedia<<ResourceOffset | MessageResponse

	//snapshot
	QuerySnapshotRequest   = OperateQuery<<OperateOffset | ResourceSnapshot<<ResourceOffset | MessageRequest
	QuerySnapshotResponse  = OperateQuery<<OperateOffset | ResourceSnapshot<<ResourceOffset | MessageResponse

	GetSnapshotRequest   = OperateGet<<OperateOffset | ResourceSnapshot<<ResourceOffset | MessageRequest
	GetSnapshotResponse  = OperateGet<<OperateOffset | ResourceSnapshot<<ResourceOffset | MessageResponse

	CreateSnapshotRequest   = OperateCreate<<OperateOffset | ResourceSnapshot<<ResourceOffset | MessageRequest
	CreateSnapshotResponse  = OperateCreate<<OperateOffset | ResourceSnapshot<<ResourceOffset | MessageResponse

	DeleteSnapshotRequest   = OperateDelete<<OperateOffset | ResourceSnapshot<<ResourceOffset | MessageRequest
	DeleteSnapshotResponse  = OperateDelete<<OperateOffset | ResourceSnapshot<<ResourceOffset | MessageResponse

	RestoreSnapshotRequest   = OperateRestore<<OperateOffset | ResourceSnapshot<<ResourceOffset | MessageRequest
	RestoreSnapshotResponse  = OperateRestore<<OperateOffset | ResourceSnapshot<<ResourceOffset | MessageResponse

	CommitSnapshotRequest   = OperateCommit<<OperateOffset | ResourceSnapshot<<ResourceOffset | MessageRequest
	CommitSnapshotResponse  = OperateCommit<<OperateOffset | ResourceSnapshot<<ResourceOffset | MessageResponse

	PullSnapshotRequest   = OperatePull<<OperateOffset | ResourceSnapshot<<ResourceOffset | MessageRequest
	PullSnapshotResponse  = OperatePull<<OperateOffset | ResourceSnapshot<<ResourceOffset | MessageResponse

	//Batch
	StartBatchCreateGuestRequest   = OperateStartBatchCreate<<OperateOffset | ResourceGuest<<ResourceOffset | MessageRequest
	StartBatchCreateGuestResponse  = OperateStartBatchCreate<<OperateOffset | ResourceGuest<<ResourceOffset | MessageResponse

	GetBatchCreateGuestRequest   = OperateGetBatchCreate<<OperateOffset | ResourceGuest<<ResourceOffset | MessageRequest
	GetBatchCreateGuestResponse  = OperateGetBatchCreate<<OperateOffset | ResourceGuest<<ResourceOffset | MessageResponse

	StartBatchDeleteGuestRequest   = OperateStartBatchDelete<<OperateOffset | ResourceGuest<<ResourceOffset | MessageRequest
	StartBatchDeleteGuestResponse  = OperateStartBatchDelete<<OperateOffset | ResourceGuest<<ResourceOffset | MessageResponse

	GetBatchDeleteGuestRequest   = OperateGetBatchDelete<<OperateOffset | ResourceGuest<<ResourceOffset | MessageRequest
	GetBatchDeleteGuestResponse  = OperateGetBatchDelete<<OperateOffset | ResourceGuest<<ResourceOffset | MessageResponse

	StartBatchStopGuestRequest  = OperateStartBatchStop<<OperateOffset | ResourceGuest<<ResourceOffset | MessageRequest
	StartBatchStopGuestResponse = OperateStartBatchStop<<OperateOffset | ResourceGuest<<ResourceOffset | MessageResponse

	GetBatchStopGuestRequest  = OperateGetBatchStop<<OperateOffset | ResourceGuest<<ResourceOffset | MessageRequest
	GetBatchStopGuestResponse = OperateGetBatchStop<<OperateOffset | ResourceGuest<<ResourceOffset | MessageResponse
	
	//Migrations
	QueryMigrationRequest   = OperateQuery<<OperateOffset | ResourceMigration<<ResourceOffset | MessageRequest
	QueryMigrationResponse  = OperateQuery<<OperateOffset | ResourceMigration<<ResourceOffset | MessageResponse

	GetMigrationRequest   = OperateGet<<OperateOffset | ResourceMigration<<ResourceOffset | MessageRequest
	GetMigrationResponse  = OperateGet<<OperateOffset | ResourceMigration<<ResourceOffset | MessageResponse

	CreateMigrationRequest   = OperateCreate<<OperateOffset | ResourceMigration<<ResourceOffset | MessageRequest
	CreateMigrationResponse  = OperateCreate<<OperateOffset | ResourceMigration<<ResourceOffset | MessageResponse

	//address pool
	QueryAddressPoolRequest  = OperateQuery<<OperateOffset | ResourceAddressPool<<ResourceOffset | MessageRequest
	QueryAddressPoolResponse = OperateQuery<<OperateOffset | ResourceAddressPool<<ResourceOffset | MessageResponse
	GetAddressPoolRequest    = OperateGet<<OperateOffset | ResourceAddressPool<<ResourceOffset | MessageRequest
	GetAddressPoolResponse   = OperateGet<<OperateOffset | ResourceAddressPool<<ResourceOffset | MessageResponse

	CreateAddressPoolRequest  = OperateCreate<<OperateOffset | ResourceAddressPool<<ResourceOffset | MessageRequest
	CreateAddressPoolResponse = OperateCreate<<OperateOffset | ResourceAddressPool<<ResourceOffset | MessageResponse
	DeleteAddressPoolRequest  = OperateDelete<<OperateOffset | ResourceAddressPool<<ResourceOffset | MessageRequest
	DeleteAddressPoolResponse = OperateDelete<<OperateOffset | ResourceAddressPool<<ResourceOffset | MessageResponse
	ModifyAddressPoolRequest  = OperateModify<<OperateOffset | ResourceAddressPool<<ResourceOffset | MessageRequest
	ModifyAddressPoolResponse = OperateModify<<OperateOffset | ResourceAddressPool<<ResourceOffset | MessageResponse

	//address range
	QueryAddressRangeRequest  = OperateQuery<<OperateOffset | ResourceAddressRange<<ResourceOffset | MessageRequest
	QueryAddressRangeResponse = OperateQuery<<OperateOffset | ResourceAddressRange<<ResourceOffset | MessageResponse
	GetAddressRangeRequest    = OperateGet<<OperateOffset | ResourceAddressRange<<ResourceOffset | MessageRequest
	GetAddressRangeResponse   = OperateGet<<OperateOffset | ResourceAddressRange<<ResourceOffset | MessageResponse

	AddAddressRangeRequest  = OperateAdd<<OperateOffset | ResourceAddressRange<<ResourceOffset | MessageRequest
	AddAddressRangeResponse = OperateAdd<<OperateOffset | ResourceAddressRange<<ResourceOffset | MessageResponse
	RemoveAddressRangeRequest  = OperateRemove<<OperateOffset | ResourceAddressRange<<ResourceOffset | MessageRequest
	RemoveAddressRangeResponse = OperateRemove<<OperateOffset | ResourceAddressRange<<ResourceOffset | MessageResponse
)

//event
const (
	EventAvailable = iota
	EventReady
	EventHeartBeat
	EventOpen
	EventClose
	EventReport
	EventCreate
	EventDelete
	EventStart
	EventStop
	EventUpdate
	EventConnect
	EventDisconnect
	EventChange
	EventAttach
	EventDetach
	EventResume
	EventAdd
	EventRemove
	EventMigrate
	EventPurge
	EventEnable
	EventDisable
	EventReset
)

const (
	ServiceAvailableEvent = EventAvailable << OperateOffset | ResourceService<<ResourceOffset | MessageEvent
	ServiceReadyEvent     = EventReady<<OperateOffset | ResourceService<<ResourceOffset | MessageEvent
	ServiceConnectedEvent = EventConnect<<OperateOffset | ResourceService<<ResourceOffset | MessageEvent
	ServiceDisconnectedEvent = EventDisconnect<<OperateOffset | ResourceService<<ResourceOffset | MessageEvent

	ConnectionOpenedEvent    = EventOpen<<OperateOffset | ResourceConnection<<ResourceOffset | MessageEvent
	ConnectionClosedEvent    = EventClose<<OperateOffset | ResourceConnection<<ResourceOffset | MessageEvent
	ConnectionKeepAliveEvent = EventHeartBeat<<OperateOffset | ResourceConnection<<ResourceOffset | MessageEvent

	CellStatusReportEvent = EventReport<<OperateOffset | ResourceComputeCell<<ResourceOffset | MessageEvent

	//core->cell
	ComputePoolReadyEvent = EventReady<<OperateOffset | ResourceComputePool<<ResourceOffset | MessageEvent

	//cell->core
	ComputeCellAvailableEvent    = EventAvailable<<OperateOffset | ResourceComputeCell<<ResourceOffset | MessageEvent
	ComputeCellReadyEvent        = EventReady<<OperateOffset | ResourceComputeCell<<ResourceOffset | MessageEvent
	ComputeCellDisconnectedEvent = EventDisconnect<<OperateOffset | ResourceComputeCell<<ResourceOffset | MessageEvent
	ComputeCellAddedEvent        = EventAdd<<OperateOffset | ResourceComputeCell<<ResourceOffset | MessageEvent
	ComputeCellRemovedEvent      = EventRemove<<OperateOffset | ResourceComputeCell<<ResourceOffset | MessageEvent

	ImageServerAvailableEvent = EventAvailable << OperateOffset | ResourceImageServer << ResourceOffset | MessageEvent

	//address pool
	AddressPoolChangedEvent = EventChange<< OperateOffset | ResourceAddressPool << ResourceOffset | MessageEvent

	//guest
	GuestCreatedEvent = EventCreate<<OperateOffset | ResourceGuest<<ResourceOffset | MessageEvent
	GuestDeletedEvent = EventDelete<<OperateOffset | ResourceGuest<<ResourceOffset | MessageEvent
	GuestStartedEvent = EventStart<<OperateOffset | ResourceGuest<<ResourceOffset | MessageEvent
	GuestStoppedEvent = EventStop<<OperateOffset | ResourceGuest<<ResourceOffset | MessageEvent
	GuestUpdatedEvent = EventUpdate<<OperateOffset | ResourceGuest<<ResourceOffset | MessageEvent
	SystemResetEvent = EventReset << OperateOffset | ResourceSystem << ResourceOffset | MessageEvent

	//instance
	InstanceMigratedEvent = EventMigrate<<OperateOffset | ResourceInstance<<ResourceOffset | MessageEvent
	InstanceAttachedEvent = EventAttach<<OperateOffset | ResourceInstance<<ResourceOffset | MessageEvent
	InstanceDetachedEvent = EventDetach<<OperateOffset | ResourceInstance<<ResourceOffset | MessageEvent
	InstancePurgedEvent   = EventPurge<<OperateOffset | ResourceInstance<<ResourceOffset | MessageEvent

	//disk image
	DiskImageUpdatedEvent = EventUpdate << OperateOffset | ResourceDiskImage << ResourceOffset | MessageEvent

	AddressChangedEvent = EventChange<< OperateOffset | ResourceAddress << ResourceOffset | MessageEvent

	//instance media
	MediaAttachedEvent = EventAttach<< OperateOffset | ResourceMedia << ResourceOffset | MessageEvent
	MediaDetachedEvent = EventDetach<< OperateOffset | ResourceMedia << ResourceOffset | MessageEvent

	//snapshot
	SnapshotResumedEvent = EventResume<< OperateOffset | ResourceSnapshot << ResourceOffset | MessageEvent
)


//ParamKey
const (
	ParamKeyName        = iota
	ParamKeyType
	ParamKeyCore
	ParamKeyMemory
	ParamKeyDisk
	ParamKeyUsage
	ParamKeyIO
	ParamKeySpeed
	ParamKeyDescription
	ParamKeyError
	ParamKeyStatus
	ParamKeyPool
	ParamKeyCell
	ParamKeyInstance
	ParamKeyEnable
	ParamKeyFlag
	ParamKeyOption
	ParamKeyID
	ParamKeyUser
	ParamKeyGroup
	ParamKeyNetwork
	ParamKeyAddress
	ParamKeyPort
	ParamKeyPriority
	ParamKeyLimit
	ParamKeyStorage
	ParamKeySource
	ParamKeyTarget
	ParamKeyMode
	ParamKeyHost
	ParamKeyPath
	ParamKeyMonitor
	ParamKeyProgress
	ParamKeyCount
	ParamKeySize
	ParamKeyAvailable
	ParamKeyMedia
	ParamKeyTag
	ParamKeyImage
	ParamKeyGuest
	ParamKeySecret
	ParamKeyImmediate
	ParamKeySystem
	ParamKeyVersion
	ParamKeyAdmin
	ParamKeyModule
	ParamKeyPrevious
	ParamKeyNext
	ParamKeyCurrent
	ParamKeyCreate
	ParamKeyModify
	ParamKeyAttach
	ParamKeyMigration
	ParamKeyStart
	ParamKeyEnd
	ParamKeyMask
	ParamKeyAllocate
	ParamKeyGateway
	ParamKeyServer
	ParamKeyAssign
	ParamKeyInternal
	ParamKeyExternal
	ParamKeyDay
	ParamKeyHour
	ParamKeyMinute
	ParamKeySecond
	ParamKeyHardware
)