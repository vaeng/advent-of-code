defmodule AdventOfCode.Day15 do
  def part1(args) do
    coords = parseToIntLists(args)
    line = 2_000_000
    # get Beacons in the line
    beaconsInLine =
      coords
      |> Enum.filter(fn [_sensorX, _sensorY, _beaconX, beaconY] -> beaconY == line end)
      |> Enum.uniq_by(fn [_sensorX, _sensorY, beaconX, _beaconY] -> beaconX end)
      |> Enum.count()

    # get coverage
    coverage =
      getCoverageForLine(coords, line, 40_000_000)
      |> MapSet.size()

    coverage - beaconsInLine
  end

  def getCoverageForLine(coords, line, xMax) do
    coords
    |> Enum.map(&coveredRangeInLine(&1, line, xMax))
    |> Enum.reduce(MapSet.new(), fn rng, acc -> MapSet.union(acc, MapSet.new(rng)) end)
  end

  def part2(args) do
    # coords = parseToIntLists(args)
    xMax = 4_000_000
    yMax = 4_000_000
    searchSpace = [0..xMax, 0..yMax]
    # points = for(x <- 0..xMax, do: for(y <- 0..yMax, do: [x, y])) |> Enum.concat()

    args
    |> String.trim()
    |> String.split("\n")
    |> Enum.map(&formatLineTo4Some/1)
    |> Enum.reduce(searchSpace, &reduceSearchSpace(&1, &2))
  end

  def parseToIntLists(args) do
    Regex.scan(~r/-?\d+/, args) |> Enum.map(&String.to_integer(hd(&1))) |> Enum.chunk_every(4)
  end

  def outOfRange?([sensorX, sensorY, beaconX, beaconY], xMax, yMax) do
    radius = abs(sensorX - beaconX) + abs(sensorY - beaconY)
    left = sensorX - radius
    right = sensorX + radius
    up = sensorY - radius
    down = sensorY + radius

    Range.disjoint?(0..xMax, left..right) or Range.disjoint?(0..yMax, up..down)
  end

  def coveredRangeInLine([sensorX, sensorY, beaconX, beaconY], line, xMax) do
    left..right = coveredRangeInLine([sensorX, sensorY, beaconX, beaconY], line)
    min(0, left)..min(xMax, right)
  end

  def coveredRangeInLine([sensorX, sensorY, beaconX, beaconY], line) do
    radius = abs(sensorX - beaconX) + abs(sensorY - beaconY)
    distanceToLine = abs(line - sensorY)
    left = sensorX - radius + distanceToLine
    right = sensorX + radius - distanceToLine

    if distanceToLine <= radius do
      min(0, left)..right
    else
      # empty range
      0..-1//1
    end
  end

  def formatLineTo4Some(line) do
    [sensorX, sensorY, beaconX, beaconY] =
      Regex.scan(~r/-?\d+/, line) |> Enum.map(&String.to_integer(hd(&1)))

    radius = abs(sensorX - beaconX) + abs(sensorY - beaconY)
    left = sensorX - radius
    right = sensorX + radius
    up = sensorY - radius
    down = sensorY + radius
    [left..right, up..down]
  end

  def reduceSearchSpace([otherl..otherr, ou..od], [l..r, u..d]) do
    cond do
      # cuts north east:
      l <= otherl and d >= od and otherl <= r and od >= u ->
        delta = max(r - otherl, d - od)
        [l..(r - delta), (u + delta)..d]

      # cuts north west:
      r >= otherr and d >= od and otherr >= l and od >= u ->
        delta = max(otherr - l, d - od)
        [(l + delta)..r, (u + delta)..d]

      # cuts south east:
      l <= otherl and u <= ou and otherl <= r and ou <= d ->
        delta = max(r - otherl, u - ou)
        [l..(r - delta), u..(d - delta)]

      # cuts south west:
      r >= otherr and u <= ou and otherr >= l and ou <= d ->
        delta = max(otherr - l, u - ou)
        [(l + delta)..r, u..(d - delta)]

      true ->
        [l..r, u..d]
    end
  end
end
