defmodule AdventOfCode.Day05 do
  def part1(args) do
    rawStack = args |> String.split("\n\n") |> hd

    indexes =
      rawStack
      |> String.split("\n")
      |> List.last()
      |> String.split()
      |> Enum.map(&String.to_integer(&1))

    stack = indexes |> Enum.map(&stackFromIndex(rawStack, &1))

    keyStack =
      Enum.zip_with([indexes, stack], fn [i, s] -> {String.to_atom(Integer.to_string(i)), s} end)

    moves = args |> String.split("\n\n") |> List.last() |> String.trim() |> String.split("\n")

    moves
    |> Enum.reduce(keyStack, fn x, acc -> makeMove(acc, x) end)
    |> getTopCrates()
  end

  @doc """
  Transform a list of lists to a String:

  Example
  iex >  getTopStacks([["D", "N", "Z"], ["C", "M"], ["P"]])
  "DCP"
  """
  def getTopCrates(keyStack),
    do:
      Keyword.values(keyStack)
      |> Enum.map(fn xs -> if(xs == [], do: "", else: hd(xs)) end)
      |> List.to_string()

  @doc """
  Transform a raw stack string to a list of  stacks, top to bottom.any()

  Example
  iex >  stackFromIndex("    [D]\n[N] [C]\n[Z] [M] [P]\n 1   2   3", 2)
  ["D", "C", "M"]

  """
  def stackFromIndex(str, idx),
    do:
      str
      |> String.split("\n")
      |> Enum.map(&String.at(&1, (idx - 1) * 4 + 1))
      |> Enum.filter(fn x -> not is_nil(x) and x != " " end)
      |> Enum.drop(-1)

  @doc """
  Transfers a list of stacks according to the move

  Example
  iex >  makeMove(["A"], ["D", "C", "M"], "move 1 from 1 to 2")
  [[], ["A", "D", "C", "M"]]
  """
  def makeMove(keyStack, move) do
    [amount, from, to] =
      move
      |> String.split()
      |> Enum.drop_every(2)
      |> Enum.map(&String.to_integer/1)

    from_atom = String.to_atom("#{from}")

    to_atom = String.to_atom("#{to}")
    crates = keyStack |> Keyword.get(from_atom) |> Enum.take(amount) |> Enum.reverse()

    keyStack
    |> Keyword.update!(to_atom, fn xs -> crates ++ xs end)
    |> Keyword.update!(from_atom, &Enum.drop(&1, amount))
  end

  def part2(args) do
    rawStack = args |> String.split("\n\n") |> hd

    indexes =
      rawStack
      |> String.split("\n")
      |> List.last()
      |> String.split()
      |> Enum.map(&String.to_integer(&1))

    stack = indexes |> Enum.map(&stackFromIndex(rawStack, &1))

    keyStack =
      Enum.zip_with([indexes, stack], fn [i, s] -> {String.to_atom(Integer.to_string(i)), s} end)

    moves = args |> String.split("\n\n") |> List.last() |> String.trim() |> String.split("\n")

    moves
    |> Enum.reduce(keyStack, fn x, acc -> makeMove9001(acc, x) end)
    |> getTopCrates()
  end

  @spec makeMove9001([{atom, any}, ...], binary) :: [{atom, any}, ...]
  def makeMove9001(keyStack, move) do
    [amount, from, to] =
      move
      |> String.split()
      |> Enum.drop_every(2)
      |> Enum.map(&String.to_integer/1)

    from_atom = String.to_atom("#{from}")

    to_atom = String.to_atom("#{to}")
    crates = keyStack |> Keyword.get(from_atom) |> Enum.take(amount)

    keyStack
    |> Keyword.update!(to_atom, fn xs -> crates ++ xs end)
    |> Keyword.update!(from_atom, &Enum.drop(&1, amount))
  end
end
