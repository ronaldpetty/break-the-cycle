module example.com

go 1.22.2

replace bs.com/pkg1 v0.0.0 => ./pkg1

replace bs.com/pkg2 v0.0.0 => ./pkg2

require bs.com/pkg2 v0.0.0

require bs.com/pkg1 v0.0.0 // indirect
