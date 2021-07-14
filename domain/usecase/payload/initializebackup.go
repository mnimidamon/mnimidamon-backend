package payload

type InitializeBackupPayload struct {
	FileName string
	Size     uint
	Hash     string

	OwnerID uint
	GroupID uint
}
