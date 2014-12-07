select id, name, isnull(parent_id, 0)
from dbo.EconomicActivities
where deleted = 0
