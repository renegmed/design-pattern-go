package observer

type Subject interface {
	/**
	 * Both of these methods take an Observer as argument
	 */
	RegisterObserver(Observer)
	DeregisterObserver(Observer)
	/**
	 * This method is called to notify all the observers
	 * when the Subject's state has changed
	 */
	NotifyObservers()
}
