defmodule DndCharacter do
  @type t :: %__MODULE__{
          strength: pos_integer(),
          dexterity: pos_integer(),
          constitution: pos_integer(),
          intelligence: pos_integer(),
          wisdom: pos_integer(),
          charisma: pos_integer(),
          hitpoints: pos_integer()
        }

  defstruct ~w[strength dexterity constitution intelligence wisdom charisma hitpoints]a

  @spec modifier(pos_integer()) :: integer()
  def modifier(score)  do
    ((score - 10) / 2) |> floor()
  end

  @spec ability :: pos_integer()
  def ability do
    1..4
    |> Enum.map(fn _ -> Enum.random(1..6) end)
    |> Enum.sort(:desc)
    |> Enum.take(3)
    |> Enum.sum()
  end

  @spec character :: t()
  def character do
    
    %__MODULE__{}
    |> Map.to_list()
    |> Enum.reduce(%__MODULE__{}, fn 
      {:hitpoints, _}, acc -> %{acc | hitpoints: 10 + modifier(acc[:constitution])}
      {k,_}, acc -> Map.put(acc, k, ability())  
    end)
  end
end
