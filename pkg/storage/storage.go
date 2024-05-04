package storage

// In my opinion, the music track mp3 file should be store on cloud storage like AWS S3, Google Cloud Storage, etc.
// The metadata of the music track should be stored in a NoSQL database like MongoDB.
// This storage package is the abstraction layer for the storage system.
// I will only implement storage in local file system, But in production, we should use cloud storage.
// This interface ensures dependency inversion principle and makes the code more testable, maintainable, and scalable.
// When using cloud storage, we can easily switch the implementation by changing the implementation of this interface.
type StorageInterface interface {
	// SaveFile saves the file to the storage system and returns the file path or URL
	SaveFile(file []byte, fileName string) (string, error)
	// GetFile gets the file from the storage system by the file path or URL
	GetFile(filePath string) ([]byte, error)
	// DeleteFile deletes the file from the storage system by the file path or URL
	DeleteFile(filePath string) error
}
