package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	ballX, ballY                float64
	ballVX, ballVY              float64
	paddleWidth, paddleHeight   float64
	leftPaddleY, rightPaddleY   float64
}

func NewGame() *Game {
	return &Game{
		ballX:          screenWidth / 2,
		ballY:          screenHeight / 2,
		ballVX:         2,
		ballVY:         2,
		paddleWidth:    10,
		paddleHeight:   80,
		leftPaddleY:    screenHeight/2 - 40,
		rightPaddleY:   screenHeight/2 - 40,
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
    return screenWidth, screenHeight
}

func (g *Game) Update() error {
	// Aggiorna la posizione della palla
	g.ballX += g.ballVX
	g.ballY += g.ballVY

	// Rimbalzo sui bordi superiore e inferiore
	if g.ballY < 0 || g.ballY > screenHeight {
		g.ballVY = -g.ballVY
	}

	// Controlla la collisione con la racchetta sinistra
	if g.ballX <= g.paddleWidth && g.ballY >= g.leftPaddleY && g.ballY <= g.leftPaddleY+g.paddleHeight {
		g.ballVX = -g.ballVX
	}

	// Controlla la collisione con la racchetta destra
	if g.ballX >= screenWidth-g.paddleWidth && g.ballY >= g.rightPaddleY && g.ballY <= g.rightPaddleY+g.paddleHeight {
		g.ballVX = -g.ballVX
	}

	// Movimento delle racchette:
	// Racchetta sinistra: W per su, S per giù
	if ebiten.IsKeyPressed(ebiten.KeyW) && g.leftPaddleY > 0 {
		g.leftPaddleY -= 4
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) && g.leftPaddleY < screenHeight-g.paddleHeight {
		g.leftPaddleY += 4
	}

	// Racchetta destra: freccia Su per su, freccia Giù per giù
	if ebiten.IsKeyPressed(ebiten.KeyUp) && g.rightPaddleY > 0 {
		g.rightPaddleY -= 4
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) && g.rightPaddleY < screenHeight-g.paddleHeight {
		g.rightPaddleY += 4
	}

	// Se la palla esce dallo schermo, la resettiamo al centro
	if g.ballX < 0 || g.ballX > screenWidth {
		g.ballX = screenWidth / 2
		g.ballY = screenHeight / 2
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Disegna la palla (un rettangolo rosso)
	ebitenutil.DrawRect(screen, g.ballX-5, g.ballY-5, 10, 10, color.RGBA{R: 0xFF, G: 0x00, B: 0x00, A: 0xFF})

	// Disegna la racchetta sinistra (rettangolo verde)
	ebitenutil.DrawRect(screen, 0, g.leftPaddleY, g.paddleWidth, g.paddleHeight, color.RGBA{R: 0x00, G: 0xFF, B: 0x00, A: 0xFF})

	// Disegna la racchetta destra (rettangolo verde)
	ebitenutil.DrawRect(screen, screenWidth-g.paddleWidth, g.rightPaddleY, g.paddleWidth, g.paddleHeight, color.RGBA{R: 0x00, G: 0xFF, B: 0x00, A: 0xFF})
}

func main() {
	game := NewGame()
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Tennis Semplice in Go")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

