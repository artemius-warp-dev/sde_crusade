defmodule CustomSet do
  defstruct map: %{}
  @opaque t :: %__MODULE__{map: map}

  @spec new(Enum.t()) :: t
  def new(enumerable) do
    %__MODULE__{map: Enum.reduce(enumerable, %{}, fn x, acc -> Map.put(acc, x, x) end)}
  end

  @spec empty?(t) :: boolean
  def empty?(custom_set) do
    custom_set.map
    |> Map.keys() == []
  end

  @spec contains?(t, any) :: boolean
  def contains?(custom_set, element) do
    Map.has_key?(custom_set.map, element)
  end

  @spec subset?(t, t) :: boolean
  # def subset?()
  def subset?(custom_set_1, custom_set_2) do
    keys1 = Map.keys(custom_set_1.map)
    keys2 = Map.keys(custom_set_2.map)
    Enum.all?(keys1, fn k1 -> k1 in keys2 end)
  end

  @spec disjoint?(t, t) :: boolean
  def disjoint?(custom_set_1, custom_set_2) do
    keys1 = Map.keys(custom_set_1.map)
    keys2 = Map.keys(custom_set_2.map)
    not Enum.any?(keys1, fn k1 -> k1 in keys2 end)
  end

  @spec equal?(t, t) :: boolean
  def equal?(custom_set_1, custom_set_2) do
    custom_set_1.map == custom_set_2.map
  end

  @spec add(t, any) :: t
  def add(custom_set, element) do
    %__MODULE__{map: Map.put(custom_set.map, element, element)}
  end

  @spec intersection(t, t) :: t
  def intersection(custom_set_1, custom_set_2) do
    keys1 = Map.keys(custom_set_1.map)
    keys2 = Map.keys(custom_set_2.map)

    Enum.filter(keys1, fn k1 -> k1 in keys2 end)
    |> new()
  end

  @spec difference(t, t) :: t
  def difference(custom_set_1, custom_set_2) do
    keys1 = Map.keys(custom_set_1.map)
    keys2 = Map.keys(custom_set_2.map)

    (keys1 -- keys2)
    |> new()
  end

  @spec union(t, t) :: t
  def union(custom_set_1, custom_set_2) do
    keys1 = Map.keys(custom_set_1.map)
    keys2 = Map.keys(custom_set_2.map)

    (keys1 ++ keys2)
    |> new()
  end
end
