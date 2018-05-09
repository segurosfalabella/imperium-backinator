package test

import "github.com/DATA-DOG/godog"

func anEndpoint() error {
	return godog.ErrPending
}

func iCallGetSnapshot() error {
	return godog.ErrPending
}

func iGetSnapshotBackupFile() error {
	return godog.ErrPending
}

func anSnapshot() error {
	return godog.ErrPending
}

func iCallCloudStorageEndpoint() error {
	return godog.ErrPending
}

func iPutBackupFileInCloudStorage() error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^an endpoint$`, anEndpoint)
	s.Step(`^I call get snapshot$`, iCallGetSnapshot)
	s.Step(`^I get snapshot backup file$`, iGetSnapshotBackupFile)
	s.Step(`^an snapshot$`, anSnapshot)
	s.Step(`^I call cloud storage endpoint$`, iCallCloudStorageEndpoint)
	s.Step(`^I put backup file in cloud storage$`, iPutBackupFileInCloudStorage)
}
