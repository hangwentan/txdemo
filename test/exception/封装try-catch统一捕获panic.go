package exception

import "fmt"

type Block struct {
	Try     func()
	Catch   func(interface{})
	Finally func()
}

func (t Block) Do() {
	if t.Finally != nil {
		defer t.Finally()
	}
	if t.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				t.Catch(r)
			}
		}()
	}
	t.Try()
}

func main() {
	Block{
		Try: func() {
			beginTransaction()
			if err = one(); err != nil {
				panic(err)
			}
			if err = two(); err != nil {
				panic(err)
			}
			if err = three(); err != nil {
				panic(err)
			}
			if err = four(); err != nil {
				panic(err)
			}
			if err = five(); err != nil {
				panic(err)
			}
			err = nil
			commit()
		},
		Catch: func(e interface{}) {
			rollback()
			fmt.Printf("%v panic\n", e)
			err = fmt.Errorf("%v", e)
		},
	}.Do()
}
