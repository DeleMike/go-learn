module example.com/go-101

go 1.23.1

replace example.com/calendar => ../calendar

replace example.com/gadget => ../gadget

require (
	example.com/calendar v0.0.0-00010101000000-000000000000
	example.com/gadget v0.0.0-00010101000000-000000000000
)
