defmodule AdventOfCode.Day23 do
  import AdventOfCode.Util

  def part1(args) do
    start_cardinals = [:north, :south, :west, :east]

    initial_map = parse_input(args)

    Enum.reduce(1..10, [start_cardinals, initial_map], fn _, [cardinals, map] ->
      next_map = round(map, cardinals)
      next_cardinals = rotate_cardinals(cardinals)
      [next_cardinals, next_map]
    end)
    |> List.last()
    |> count_empty_tiles()
  end

  def part2(args) do
    start_cardinals = [:north, :south, :west, :east]

    initial_map = parse_input(args)

    Enum.reduce_while(1..99999, [start_cardinals, initial_map], fn round, [cardinals, map] ->
      next_map = round(map, cardinals)
      next_cardinals = rotate_cardinals(cardinals)
      if next_map == map, do: {:halt, round}, else: {:cont, [next_cardinals, next_map]}

    end)


  end

  def count_empty_tiles(map) do
    {[x_min, _], [x_max, _]} = Map.keys(map) |> Enum.min_max_by(fn [x,_y] -> x end)
    {[_, y_min], [_, y_max]} = Map.keys(map) |> Enum.min_max_by(fn [_x,y] -> y end)

    elves = Map.size(map)
    (x_max - x_min + 1) * (y_max - y_min + 1) - elves
  end

  def rotate_cardinals([a, b, c, d]), do: [b, c, d, a]

  def round(map, ordered_cadinals) do
    map
    |> get_proposals(ordered_cadinals)
    |> filter_collisions()
  end

  def print(map) do
    {[x_min, _], [x_max, _]} = Map.keys(map) |> Enum.min_max_by(fn [x,_y] -> x end)
    {[_, y_min], [_, y_max]} = Map.keys(map) |> Enum.min_max_by(fn [_x,y] -> y end)

    y_min..y_max
    |> Enum.map(&(Enum.map(x_min..x_max,
        fn x ->
          Map.get(map, [x, &1], ".")
      end)))
    |> Enum.intersperse(["\n"])
    |> IO.puts()

    IO.puts("")
    map

  end

  def parse_input(args) do
    args
    |> String.split("\n")
    |> Enum.with_index(fn line, row ->
      Enum.with_index(String.codepoints(line), fn letter, col -> if letter == "#", do: [[col, row], letter], else: [] end)
    end)
    |> List.flatten()
    |> Enum.chunk_every(3)
    |> Enum.reduce(Map.new(), fn [x, y, k], map -> Map.put(map, [x, y], k) end)
  end

  def get_proposals(map, ordered_cardinals) do
    Map.keys(map)
    |> Enum.map(fn pos ->
      [:nothing | ordered_cardinals]
      |> Enum.map(fn card -> propose_direction(pos, map, card)
      end)
      |> Enum.filter(&(&1))
      |> (then &(if &1 == [], do: pos, else: hd(&1)))
      |> (then &({pos, &1}))
     end)
  end

  def filter_collisions(proposals) do
    good_proposals = proposals
    |> Enum.frequencies_by(fn {_from, to} -> to end)
    |> Enum.reject(fn {_k, v} -> v > 1 end)
    |> Enum.map(fn {k, _v} -> k end)


    proposals
    |> Enum.map(fn {from, to} -> if Enum.member?(good_proposals, to), do: {to, "#"}, else: {from, "#"} end)
    |> Enum.into(Map.new())
  end

  def propose_direction(position, map, :nothing) do
    [x, y] = position
    [getN(position), getW(position), getS(position), getE(position), getNW(position), getNE(position), getSE(position), getSW(position)]
    |> Enum.all?(fn proposal -> Map.get(map, proposal, ".") == "." end)
    #|> IO.inspect(label: "Do nothing? #{x} #{y}")
    |> (then &(if &1, do: position, else: false))
  end

  def propose_direction(position, map, :south) do
    [x, y] = position
    [getNE(position), getNW(position), getN(position)]
    |> Enum.all?(fn proposal -> Map.get(map, proposal, ".") == "." end)
    #|> IO.inspect(label: "Go north? #{x} #{y}")
    |> (then &(if &1, do: getN(position), else: false))
  end

  def propose_direction(position, map, :east) do
    [getE(position), getSE(position), getNE(position)]
    |> Enum.all?(fn proposal -> Map.get(map, proposal, ".") == "." end)
    |> (then &(if &1, do: getE(position), else: false))
  end

  def propose_direction(position, map, :north) do
    [x, y] = position
    [getS(position), getSE(position), getSW(position)]
    |> Enum.all?(fn proposal -> Map.get(map, proposal, ".") == "." end)
    #|> IO.inspect(label: "Go south? #{x} #{y}")
    |> (then &(if &1, do: getS(position), else: false))
  end

  def propose_direction(position, map, :west) do
    [getNW(position), getW(position), getSW(position)]
    |> Enum.all?(fn proposal -> Map.get(map, proposal, ".") == "." end)
    |> (then &(if &1, do: getW(position), else: false))
  end

end
