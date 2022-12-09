defmodule AdventOfCode.Day08 do
  def part1(_args) do
  end

  def part2(_args) do
  end

  def inputToNumGrid(input) do
    input
    |> String.trim()
    |> String.split("\n")
    |> Enum.map(
      &(String.codepoints(&1)
        |> Enum.map(fn x -> String.to_integer(x) end))
    )
  end

  def setAllToUnknown(numgrid) do
    numgrid |> Enum.map(&Enum.map(&1, fn x -> [x, :unknown] end))
  end

  def initializeBoarders(tagGrid) do
    tagGrid
    |> List.update_at(-1, &Enum.map(&1, fn [x, _] -> [x, :visible] end))
    |> List.update_at(0, &Enum.map(&1, fn [x, _] -> [x, :visible] end))
    |> Enum.map(fn [[x, _] | rest] -> [[x, :visible] | rest] end)
    |> Enum.map(&List.update_at(&1, -1, fn [x, _] -> [x, :visible] end))
  end

  def updateStatus(_, _, _, :visible, _), do: :visible
  def updateStatus(_, _, _, :invisible, _), do: :invisible

  def updateStatus(rix, cix, height, :unknown, grid) do
    # [[hup, sup], [hdown, sdown], [hleft, sleft], [hright, sright]] = getEnvironment(grid, rix, cix)
    getEnvironment(grid, rix, cix) |> Enum.any?()

    # cases for: all blocking -> invis
    # any is smaller and vis? -> vis
    # else -> unknown
  end

  def updateStatus(_, _, _, _, _), do: :error

  def updateGrid(tagGrid) do
    newStatus = fn _, _, _, _ -> :test end

    workList =
      tagGrid
      |> Enum.with_index()
      |> Enum.map(fn {row, rix} ->
        Enum.with_index(row, fn element, cix -> [rix, cix, element] end)
      end)
      |> List.flatten()
      |> Enum.chunk_every(4)

    workList
    |> Enum.reduce(tagGrid, fn [rix, cix, height, status], accgrid ->
      List.update_at(
        accgrid,
        rix,
        &List.update_at(&1, cix, fn [x, status] ->
          [x, updateStatus(rix, cix, height, status, accgrid)]
        end)
      )
    end)
  end

  def getVisibleTrees(grid), do: grid |> List.flatten() |> Enum.count(&(&1 == :visible))

  def updateStatusByCoordinate(grid, rix, cix, newStatus),
    do: grid |> List.update_at(rix, &List.update_at(&1, cix, fn [x, _] -> [x, newStatus] end))

  def getElement(grid, rix, cix),
    do: grid |> Enum.fetch(rix) |> then(fn {_, row} -> Enum.fetch!(row, cix) end)

  def(getEnvironment(grid, rix, cix)) do
    up = getElement(grid, rix - 1, cix)
    down = getElement(grid, rix + 1, cix)
    left = getElement(grid, rix, cix - 1)
    right = getElement(grid, rix, cix + 1)
    [up, down, left, right]
  end
end
