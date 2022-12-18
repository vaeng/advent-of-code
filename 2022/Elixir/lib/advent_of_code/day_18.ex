defmodule AdventOfCode.Day18 do
  def part1(args) do
    cubes =
      args
      |> String.trim()
      |> String.split("\n")
      |> Enum.map(&(String.split(&1, ",") |> Enum.map(fn x -> String.to_integer(x) end)))

    max = Enum.count(cubes) * 6
    xy = Enum.map(cubes, fn [x, y, z] -> [x, y] end) |> Enum.uniq() |> Enum.count()
    yz = Enum.map(cubes, fn [x, y, z] -> [y, z] end) |> Enum.uniq() |> Enum.count()
    zx = Enum.map(cubes, fn [x, y, z] -> [z, x] end) |> Enum.uniq() |> Enum.count()

    visible = 2 * (xy + yz + zx)
    visible + (max - visible) / 2
  end

  def part2(_args) do
  end
end
