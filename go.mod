module example.com

go 1.22.2

replace bs.com/pkg1 v0.0.0 => ./pkg1

replace bs.com/pkg2 v0.0.0 => ./pkg2

replace bs.com/pkg3 v0.0.0 => ./pkg3

replace bs.com/pkg4 v0.0.0 => ./pkg4

require (
	bs.com/pkg2 v0.0.0 // indirect
	bs.com/pkg4 v0.0.0 // indirect
)

require (
	bs.com/pkg1 v0.0.0
	bs.com/pkg3 v0.0.0 // indirect
)
