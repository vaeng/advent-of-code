mkdir 01
mkdir 02
mkdir 03
mkdir 04
mkdir 05
mkdir 06
mkdir 07
mkdir 08
mkdir 09
mkdir 10
mkdir 11
mkdir 12
mkdir 13
mkdir 14
mkdir 15
mkdir 16
mkdir 17
mkdir 18
mkdir 19
mkdir 20
mkdir 21
mkdir 22

mklink .\03\03.go ..\..\..\2019\Clojure\day3\src\day3\core.clj

for /l %x in (1, 1, 9) do (
   echo %x
   mklink .\0%x\0%x.go ..\..\..\2021\Go\day0%x\day0%x.go
)

for /l %x in (0, 1, 5) do (
   echo %x
   mklink .\1%x\1%x.go ..\..\..\2021\Go\day1%x\day1%x.go
)

for /l %x in (6, 1, 9) do (
   echo %x
   mklink .\1%x\1%x.go ..\..\..\2021\Go\day1%x\main.go
)

for /l %x in (0, 1, 2) do (
   echo %x
   mklink .\2%x\2%x.go ..\..\..\2021\Go\day2%x\main.go
)