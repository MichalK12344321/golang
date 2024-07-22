package collector

type CollectionJobData struct {
	stdout chan []byte
	stderr chan []byte
	errors chan error
	done   chan any
}

func (job *CollectionJobData) StdoutChannel() chan []byte {
	return job.stdout
}

func (job *CollectionJobData) StderrChannel() chan []byte {
	return job.stderr
}

func (job *CollectionJobData) ErrorChannel() chan error {
	return job.errors
}

func (job *CollectionJobData) DoneChannel() chan any {
	return job.done
}

func NewCollectionJobData() *CollectionJobData {
	return &CollectionJobData{
		stdout: make(chan []byte),
		stderr: make(chan []byte),
		errors: make(chan error, 1),
		done:   make(chan any, 1),
	}
}
