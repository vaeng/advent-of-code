defmodule AdventOfCode.Day18 do
  import AdventOfCode.Util

  def part1(args) do
    cubes = parseInput(args)

    xy =
      cubes
      |> then(&countGroupsOf3dDimenion(&1))

    yz =
      cubes
      |> Enum.map(fn [x, y, z] -> [y, z, x] end)
      |> then(&countGroupsOf3dDimenion(&1))

    zx =
      cubes
      |> Enum.map(fn [x, y, z] -> [z, x, y] end)
      |> then(&countGroupsOf3dDimenion(&1))

    2 * (xy + yz + zx)
  end

  def parseInput(args) do
    args
    |> String.trim()
    |> String.split("\n")
    |> Enum.map(&(String.split(&1, ",") |> Enum.map(fn x -> String.to_integer(x) end)))
  end

  def countGroupsOf3dDimenion(cubes) do
    cubes
    |> Enum.sort()
    |> Enum.chunk_by(fn [x, y, _z] -> [x, y] end)
    |> Enum.map(fn x ->
      Enum.map(
        x,
        &Enum.at(&1, 2)
      )
      |> then(
        &(Enum.scan(&1, [hd(&1), true], fn x, acc -> [x, hd(acc) + 1 == x] end)
          |> Enum.count(fn [_, x] -> x == false end))
      )
    end)
    |> Enum.sum()
  end

  def part2(args) do
    cubes = Enum.reduce(parseInput(args), MapSet.new(), &MapSet.put(&2, &1))

    {xMin, xMax} = Enum.map(cubes, fn [x, _y, _z] -> x end) |> Enum.min_max()
    {yMin, yMax} = Enum.map(cubes, fn [_x, y, _z] -> y end) |> Enum.min_max()
    {zMin, zMax} = Enum.map(cubes, fn [_x, _y, z] -> z end) |> Enum.min_max()

    startCube = [xMin - 1, yMin, zMin]
    discoveredCubes = MapSet.new()
    queue = [startCube]

    surroundingWater = bfs3d(discoveredCubes, queue, cubes, xMin, yMin, zMin, xMax, yMax, zMax)

    northFaces =
      Enum.map(surroundingWater, fn cube -> MapSet.member?(cubes, getNext(cube, :north)) end)
      |> Enum.filter(& &1)
      |> Enum.count()

    westFaces =
      Enum.map(surroundingWater, fn cube -> MapSet.member?(cubes, getNext(cube, :west)) end)
      |> Enum.filter(& &1)
      |> Enum.count()

    southFaces =
      Enum.map(surroundingWater, fn cube -> MapSet.member?(cubes, getNext(cube, :south)) end)
      |> Enum.filter(& &1)
      |> Enum.count()

    eastFaces =
      Enum.map(surroundingWater, fn cube -> MapSet.member?(cubes, getNext(cube, :east)) end)
      |> Enum.filter(& &1)
      |> Enum.count()

    upFaces =
      Enum.map(surroundingWater, fn cube -> MapSet.member?(cubes, getNext(cube, :up)) end)
      |> Enum.filter(& &1)
      |> Enum.count()

    downFaces =
      Enum.map(surroundingWater, fn cube -> MapSet.member?(cubes, getNext(cube, :down)) end)
      |> Enum.filter(& &1)
      |> Enum.count()

    [northFaces, westFaces, southFaces, eastFaces, upFaces, downFaces] |> Enum.sum()
  end

  def reachable?([x, y, z], cubes, xMin, yMin, zMin, xMax, yMax, zMax) do
    cond do
      x < xMin - 1 -> false
      x > xMax + 1 -> false
      y < yMin - 1 -> false
      y > yMax + 1 -> false
      z < zMin - 1 -> false
      z > zMax + 1 -> false
      MapSet.member?(cubes, [x, y, z]) -> false
      true -> true
    end
  end

  def bfs3d(discoveredCubes, queue, cubes, xMin, yMin, zMin, xMax, yMax, zMax) do
    if queue == [] do
      discoveredCubes
    else
      [h | rest] = queue

      newElements =
        Enum.map([:north, :east, :west, :south, :up, :down], &getNext(h, &1))
        |> Enum.filter(&reachable?(&1, cubes, xMin, yMin, zMin, xMax, yMax, zMax))
        |> Enum.reject(&MapSet.member?(discoveredCubes, &1))

      bfs3d(
        MapSet.put(discoveredCubes, h),
        Enum.uniq(rest ++ newElements),
        cubes,
        xMin,
        yMin,
        zMin,
        xMax,
        yMax,
        zMax
      )
    end
  end
end
