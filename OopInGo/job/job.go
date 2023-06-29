package job

type Job struct {
	Name     string
	Location string
	Remote   bool
}

func (j *Job) SetJobName(name string) {
	j.Name = name
}
