package actors

type ActorInfo struct {
	identifier string
	ch         chan Event
}

func makeActorInfo(identifier string, ch chan Event) *ActorInfo {
	actorInfo := &ActorInfo{
		identifier: identifier,
		ch:         ch,
	}

	return actorInfo
}

func (info *ActorInfo) Identifier() string {
	return info.identifier
}
