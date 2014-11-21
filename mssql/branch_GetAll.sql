select id,
    isnull(name, '') name,
    isnull(code, '') code,
    isnull(description, '') description,
    isnull(address, '') address
from dbo.Branches
where deleted = 0
