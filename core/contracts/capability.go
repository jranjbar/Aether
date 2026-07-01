package contracts

type Capability interface{Name() string; Execute(ctx any) error}
