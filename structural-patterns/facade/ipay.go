package facade

type IPay interface {
	Pay(cost float64) bool
}
