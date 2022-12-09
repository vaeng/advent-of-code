defmodule AdventOfCode.Day082015 do
  def part1(_args) do
  end

  def part2(_args) do
  end

  def numberOfCodeChars(_input) do
    {:ok, content} = File.read("./lib/advent_of_code/day_08.input")
    content |> String.split("\n") |> Enum.map(&String.length(&1))
  end

  def numberOfMemoryChars(_input) do
    {:ok, content} = File.read("./lib/advent_of_code/day_08.input")

    content |> String.split("\n")    |> Enum.map(&String.slice(&1, 1..-2))    |> Enum.map(&String.replace(&1, ~r/\\x\d\d/, "A"))    |> Enum.map(&String.replace(&1, ~r/\\\\/, "A"))    |> Enum.map(&String.replace(&1, ~r/\\\"/, "A"))|> Enum.map(&String.length/1)
  end
end
