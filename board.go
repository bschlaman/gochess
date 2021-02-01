package board

type Side int

const (
	White = iota
	Black
)

func (s Side) String() string {
	return [...]string{"white", "black"}[s]
}

type Move struct {
	To    int
	From  int
	Flags int
}

type Irrev struct {
	move          Move
	EnPasSq       int
	CPerm         int
	CapturedPiece int
	Pinned        uint64
}

type BoardState struct {
	Board      [120]int
	Ply        int
	SideToMove bool
	EnPasSq    int
	// 0 0 0 0
	CPerm   int
	KingSq  [2]int
	Pinned  uint64
	History [200]Irrev
}

func NewBoardState() *BoardState {
	var BS BoardState
	return &BS
}

type Type int

const (
	Knight Type = iota
	Bishop
	Rook
	Queen
	King
)

const (
	NORMAL_MODE = iota
	FEN_MODE
	PERFT_MODE
	O_64_BOARD
	O_BOARD_STATE
	O_120_BOARD
	O_PINNED
)
