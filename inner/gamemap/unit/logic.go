package unit

import (
	"fmt"
	"log"

	"github.com/Frosin/Communis/inner/consts"
	"github.com/Frosin/Communis/inner/limits"
	"github.com/Frosin/Communis/inner/random"
)

type target struct {
	tX, tY int
	headsOrTails,
	escapeX,
	escapeY bool
}

type unitLogic struct {
	targets    []target
	limits     *limits.Limits
	isIdle     bool
	navigation uint8
}

var (
	escapeStep = 5
)

func newLogic(limits *limits.Limits) unitLogic {
	return unitLogic{
		limits: limits,
	}
}

func (u *unitLogic) SetTarget(tX, tY int) {
	u.targets = append(u.targets, target{tX, tY, random.HeadsOrTails(), false, false})
	u.isIdle = false
}

func (u *unitLogic) getNavigationX() bool {
	return 0 == consts.Left&u.navigation
}

func (u *unitLogic) getNavigationY() bool {
	return 0 == consts.Up&u.navigation
}

func (u *unitLogic) setNavigation(newNav uint8) {
	switch newNav {
	case consts.Up:
		u.navigation &^= consts.Down
		u.navigation |= consts.Up
	case consts.Down:
		u.navigation &^= consts.Up
		u.navigation |= consts.Down
	case consts.Left:
		u.navigation &^= consts.Right
		u.navigation |= consts.Left
	case consts.Right:
		u.navigation &^= consts.Left
		u.navigation |= consts.Right
	}
}

func (u *unitLogic) detectBigDeadlock() {
	if len(u.targets) > 5 {
		//debug
		fmt.Println("targets:", u.targets)
		panic("deadlock!")
	}
}

func (u *unitLogic) NextXY(curX, curY int) (int, int) {
	if len(u.targets) == 0 {
		log.Println("no targets!")
		return curX, curY
	}
	u.detectBigDeadlock()

	newX, newY := curX, curY

	if u.targets[0].tY > curY && u.limits.IsValidPosition(curX, curY+1) {
		newY++
		//set down direction
		u.setNavigation(consts.Down)
	}
	if u.targets[0].tY < curY && u.limits.IsValidPosition(curX, curY-1) {
		newY--
		u.setNavigation(consts.Up)
	}
	if u.targets[0].tX > curX && u.limits.IsValidPosition(curX+1, curY) {
		newX++
		u.setNavigation(consts.Right)
	}
	if u.targets[0].tX < curX && u.limits.IsValidPosition(curX-1, curY) {
		newX--
		u.setNavigation(consts.Left)
	}
	//finish moving
	if curX == u.targets[0].tX && curY == u.targets[0].tY {
		u.isIdle = true
		//next target
		//if it is escape by X target
		if u.targets[0].escapeX {
			//if escape didn't finish
			//add new target for continue escape
			var newTarget target
			if u.targets[1].tY > curY {
				if !u.limits.IsValidPosition(curX, curY+1) {
					if u.targets[0].headsOrTails {
						newTarget = target{curX + escapeStep, u.targets[0].tY, true, true, false}
					} else {
						newTarget = target{curX - escapeStep, u.targets[0].tY, false, true, false}
					}
					u.targets[0] = newTarget
				} else {
					u.targets = u.targets[1:]
					newY++
				}
			} else {
				if !u.limits.IsValidPosition(curX, curY-1) {
					if u.targets[0].headsOrTails {
						newTarget = target{curX + escapeStep, u.targets[0].tY, true, true, false}
					} else {
						newTarget = target{curX - escapeStep, u.targets[0].tY, false, true, false}
					}
					u.targets[0] = newTarget
				} else {
					u.targets = u.targets[1:]
					newY--
				}
			}
			//if escape by Y
		} else if u.targets[0].escapeY {
			//if escape didn't finish
			//add new target for continue escape
			var newTarget target
			if u.targets[1].tX > curX {
				if !u.limits.IsValidPosition(curX+1, curY) {
					if u.targets[0].headsOrTails {
						newTarget = target{u.targets[0].tX, curY + escapeStep, true, false, true}
					} else {
						newTarget = target{u.targets[0].tX, curY - escapeStep, false, false, true}
					}
					u.targets[0] = newTarget
				} else {
					u.targets = u.targets[1:]
					newX++
				}
			} else {
				if !u.limits.IsValidPosition(curX-1, curY) {
					if u.targets[0].headsOrTails {
						newTarget = target{u.targets[0].tX, curY + escapeStep, true, false, true}
					} else {
						newTarget = target{u.targets[0].tX, curY - escapeStep, false, false, true}
					}
					u.targets[0] = newTarget
				} else {
					u.targets = u.targets[1:]
					newX--
				}
			}
		} else {
			if len(u.targets) == 1 {
				u.targets = nil
			} else {
				//remove escape terget
				u.targets = u.targets[1:]
			}
		}
		return newX, newY
	}

	//if no changes, it is deadlock
	//we will set terget for escape
	if newX == curX && newY == curY {
		println("deadlock! ", u.targets[0].tX, u.targets[0].tY, curX, curY)
		//if deadlock on Y
		if u.targets[0].tX == curX {
			var newTarget target
			if u.getNavigationX() {
				newTarget = target{newX + escapeStep, newY, true, true, false}
			} else {
				newTarget = target{newX - escapeStep, newY, false, true, false}
			}
			u.targets = append([]target{newTarget}, u.targets...)
		} else {
			//deadlock on X
			var newTarget target
			if u.getNavigationY() {
				newTarget = target{newX, newY + escapeStep, true, false, true}
			} else {
				newTarget = target{newX, newY - escapeStep, false, false, true}
			}
			u.targets = append([]target{newTarget}, u.targets...)
		}
	}

	return newX, newY
}

func (u *unitLogic) haveTargets() bool {
	return len(u.targets) != 0
}
