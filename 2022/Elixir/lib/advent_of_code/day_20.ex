defmodule AdventOfCode.Day20 do
  def part1(args) do
    dectrypt(args, 1, 1)
  end

  def part2(args) do
    dectrypt(args, 10, 811589153)
  end

  def dectrypt(args, cycles, decryptKey) do
    initialArray = args |> String.split() |> Enum.with_index(fn el, idx -> {idx, String.to_integer(el) * decryptKey } end)
    length = Enum.count(initialArray)
    mixed = 0..length-1
    |> Stream.cycle()
    |> Enum.take(cycles*length)
    |> Enum.reduce(initialArray, fn idx, arr ->
        current_pos = Enum.find_index(arr, fn {i, _el} -> i == idx end)
        current_element = Enum.at(arr, current_pos) |> elem(1)
        if current_element == 0 do
          arr
        else
          raw_pos = current_pos + current_element
          intended_pos = Integer.mod(raw_pos, length-1)
          List.delete_at(arr, current_pos) |> List.insert_at(intended_pos, {idx, current_element})
        end
      end)
    shift = Enum.find_index(mixed, fn {_i, el} -> el == 0 end)
    [1000, 2000, 3000]
    |> Enum.map(&elem(Enum.at(mixed, Integer.mod(shift+&1, length)), 1))
    |> Enum.sum()
  end
end
