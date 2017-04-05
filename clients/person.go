package clients

type Person struct {
	Id string
	NickName string
	Sex int //0 for female, 1 for male
	Age int
	Address string
	Avatar string

}

func (p *Person) SetId(id string) {
	p.Id = id
}

func (p *Person) SetSex(sex int) {
	p.Sex = sex
}

func (p *Person) SetAge(age int) {
	p.Age = age
}

func (p *Person) SetAddress(addr string) {
	p.Address = addr
}

func (p *Person) SetAvatar(url string) {
	p.Avatar = url
}
