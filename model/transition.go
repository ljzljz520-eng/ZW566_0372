package model

func CanTransition(from, to string) bool {
	switch from {
	case StatusOpen:
		return to == StatusAssigned
	case StatusAssigned:
		return to == StatusWorking
	case StatusWorking:
		return to == StatusClosed
	case StatusClosed:
		return to == StatusArchived
	case StatusArchived:
		return false
	}
	return false
}
func Statuses() []string {
	return []string{StatusOpen, StatusAssigned, StatusWorking, StatusClosed, StatusArchived}
}
func IsTerminal(s string) bool { return s == StatusClosed || s == StatusArchived }
