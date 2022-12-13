defmodule AdventOfCode.Day11 do
  def part1(args) do
    args
    |> parseInput()
    |> then(&Enum.reduce(1..10000, &1, fn _, map -> playRound(map) end))
    |> Map.values()
    |> Enum.map(& &1.counter)
    |> IO.inspect(charlists: :as_lists)
    |> Enum.sort(:desc)
    |> Enum.take(2)
    |> Enum.product()
  end

  def part2(_args) do
  end

  def parseInput(args) do
    re =
      ~r/Monkey (?<id>\d+):\n  Starting items: (?<items>\d+(, \d+)*)\n  Operation: new = old (?<op>. [\d\w]+)\n  Test: divisible by (?<test>\d+)\n    If true: throw to monkey (?<trMonk>\d+)\n    If false: throw to monkey (?<faMonk>\d+)/

    args
    |> String.split("\n\n")
    |> Enum.map(&Regex.named_captures(re, &1))
    |> Enum.reduce(%{}, &reMapToMonkeyMap(&1, &2))
  end

  def reMapToMonkeyMap(reMap, targetMap) do
    id = reMap["id"]
    items = reMap["items"] |> String.split(", ") |> Enum.map(&String.to_integer/1)
    faMonk = reMap["faMonk"]
    trMonk = reMap["trMonk"]
    test = String.to_integer(reMap["test"])

    op =
      reMap["op"]
      |> String.split()
      |> then(
        &case &1 do
          ["*", "old"] -> fn x -> x * x end
          ["+", "old"] -> fn x -> x + x end
          ["*", num] -> fn x -> x * String.to_integer(num) end
          ["+", num] -> fn x -> x + String.to_integer(num) end
        end
      )

    Map.put_new(targetMap, id, %{
      items: items,
      faMonk: faMonk,
      trMonk: trMonk,
      test: test,
      op: op,
      counter: 0
    })
  end

  def processItem(map, item, id) do
    worry = map[id].op.(item)

    afterInspect = Integer.mod(worry, 9_699_690)

    check = Integer.mod(afterInspect, map[id].test) == 0
    target = if check, do: map[id].trMonk, else: map[id].faMonk

    map
    |> update_in([target, :items], fn xs -> xs ++ [afterInspect] end)
    |> update_in([id, :items], fn xs -> tl(xs) end)
    |> update_in([id, :counter], fn i -> i + 1 end)
  end

  def playRound(monkeyMap) do
    monkeyMap
    |> Map.keys()
    |> Enum.reduce(monkeyMap, fn monkeyID, map -> playTurn(monkeyID, map) end)
  end

  def playTurn(id, map) do
    map[id].items
    |> Enum.reduce(map, fn item, accMap -> processItem(accMap, item, id) end)
  end
end
