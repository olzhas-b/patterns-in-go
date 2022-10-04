package prototype

func (edu *Education) Copy() Education {
	coursesClone := make([]string, len(edu.Courses))
	copy(coursesClone, edu.Courses)

	return Education{
		Name:    edu.Name,
		Courses: coursesClone,
	}
}

type Education struct {
	Name    string
	Courses []string
}
