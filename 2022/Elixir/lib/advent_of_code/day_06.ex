defmodule AdventOfCode.Day06 do
  def part1(args) do

    solution = args |> String.to_charlist() |> Stream.chunk_every(4,1, :discard) |> Enum.take_while(fn x -> MapSet.new(x) |> MapSet.size() != 4 end) |> Enum.count()
    solution + 4
  end

  def part2(args) do
    solution = args |> String.to_charlist() |> Stream.chunk_every(14,1, :discard) |> Enum.take_while(fn x -> MapSet.new(x) |> MapSet.size() != 14 end) |> Enum.count()
    solution + 14
  end
end
