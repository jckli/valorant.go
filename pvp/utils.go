package pvp

type PvpOptionalParams struct {
	startIndex int
	endIndex   int
	size       int
	queue      string
}

func WithStartIndex(startIndex int) func(*PvpOptionalParams) {
	return func(m *PvpOptionalParams) {
		m.startIndex = startIndex
	}
}

func WithEndIndex(endIndex int) func(*PvpOptionalParams) {
	return func(m *PvpOptionalParams) {
		m.endIndex = endIndex
	}
}

func WithQueue(queue string) func(*PvpOptionalParams) {
	return func(m *PvpOptionalParams) {
		m.queue = queue
	}
}

func WithSize(size int) func(*PvpOptionalParams) {
	return func(m *PvpOptionalParams) {
		m.size = size
	}
}
