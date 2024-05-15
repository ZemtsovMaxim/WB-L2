package pattern

/*
	Реализовать паттерн «состояние».
	Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern

	Состояние — это поведенческий паттерн проектирования, который позволяет объектам менять поведение в зависимости от своего состояния.
	Извне создаётся впечатление, что изменился класс объекта.
	Основная идея в том, что программа может находиться в одном из нескольких состояний, которые всё время сменяют друг друга.
	Набор этих состояний, а также переходов между ними, предопределён и конечен.
	Находясь в разных состояниях, программа может по-разному реагировать на одни и те же события, которые происходят с ней.

	Паттерн Состояние предлагает создать отдельные классы для каждого состояния, в котором может пребывать объект, а затем вынести туда поведения, соответствующие этим состояниям.
	Вместо того, чтобы хранить код всех состояний, первоначальный объект, называемый контекстом,
 	будет содержать ссылку на один из объектов-состояний и делегировать ему работу, зависящую от состояния.

	+
	Избавляет от множества больших условных операторов машины состояний.
 	Концентрирует в одном месте код, связанный с определённым состоянием.
 	Упрощает код контекста.

	-
	Может неоправданно усложнить код, если состояний мало и они редко меняются.
*/

import "fmt"

// Context представляет контекст, у которого есть текущее состояние
type Context struct {
	state State
}

func NewContext(state State) *Context {
	return &Context{state: state}
}

// State определяет интерфейс для конкретных состояний
type State interface {
	Handle()
}

// ConcreteStateA и ConcreteStateB - конкретные состояния, реализующие интерфейс State
type ConcreteStateA struct{}

func (csa *ConcreteStateA) Handle() {
	fmt.Println("ConcreteStateA обрабатывает запрос")
}

type ConcreteStateB struct{}

func (csb *ConcreteStateB) Handle() {
	fmt.Println("ConcreteStateB обрабатывает запрос")
}

// Пример использования

func ExampleUsage() {
	// Создаем контекст с начальным состоянием ConcreteStateA
	context := NewContext(&ConcreteStateA{})

	// Вызываем метод Handle() текущего состояния
	context.state.Handle()

	// Меняем состояние на ConcreteStateB и вызываем метод Handle() снова
	context.state = &ConcreteStateB{}
	context.state.Handle()
}
