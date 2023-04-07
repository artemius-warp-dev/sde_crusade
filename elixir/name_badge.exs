defmodule NameBadge do
  def print(id, name, department) do
    id_str = if id, do: "[#{id}]", else: ""
    department_str = if department, do: "- " <> String.upcase(department), else: "- OWNER"
    name_str = if name, do: " - #{name} ", else: ""
    id_str <> name_str <> department_str
  end
end
