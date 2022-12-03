defmodule AdventOfCode.Day03 do
  def part1(args) do
    args
    |> String.split()
    |> Enum.map(&splitInHalf/1)
    |> Enum.map(&findDuplicate/1)
    |> Enum.map(&convertToPriority/1)
    |> Enum.sum()
  end

  def splitInHalf(str), do: String.split_at(str, div(String.length(str), 2))

  def findDuplicate({str1, str2}),
    do:
      MapSet.intersection(
        MapSet.new(String.codepoints(str1)),
        MapSet.new(String.codepoints(str2))
      )
      |> MapSet.to_list()
      |> hd()

  def findCommon([str1, str2, str3]),
    do:
      MapSet.intersection(
        MapSet.new(String.codepoints(str1)),
        MapSet.intersection(
          MapSet.new(String.codepoints(str2)),
          MapSet.new(String.codepoints(str3))
        )
      )
      |> MapSet.to_list()
      |> hd()

  def convertToPriority(<<c::utf8, _::binary>>) when c <= ?z and c >= ?a, do: c - ?a + 1
  def convertToPriority(<<c::utf8, _::binary>>) when c <= ?Z and c >= ?A, do: c - ?A + 27
  def convertToPriority(_), do: :error

  def part2(args) do
    args
    |> String.split()
    |> Stream.chunk_every(3)
    |> Enum.map(&findCommon/1)
    |> Enum.map(&convertToPriority/1)
    |> Enum.sum()
  end
end
