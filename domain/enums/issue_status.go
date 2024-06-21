package enums

// const To_Do = "To Do"
const Open = "Open"
const In_Progress = "In Progress"
const Review = "Review"
const Closed = "Closed"

func GetAllStatus() []string {
	return []string{
		// To_Do,
		Open,
		In_Progress,
		Review,
		Closed,
	}
}
