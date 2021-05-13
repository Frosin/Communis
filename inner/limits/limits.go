package limits

type limitParam struct {
	x, y, w, h int
}

type Limits struct {
	limits []limitParam
}

func NewLimits() *Limits {
	return &Limits{}
}

func (l *Limits) IsValidPosition(uX, uY int) bool {
	for _, l := range l.limits {
		if uX >= l.x &&
			uX <= l.x+l.w &&
			uY >= l.y &&
			uY <= l.y+l.h {
			return false
		}
	}
	return true
}

func (l *Limits) SetLimit(rectX, rectY, rectWidth, rectHeight int) {
	l.limits = append(l.limits, limitParam{rectX, rectY, rectWidth, rectHeight})
}
