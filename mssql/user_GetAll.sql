select id, user_name, first_name, last_name 
from (
	select id, user_name, first_name, last_name, 
		row_number() over (order by id asc) num
	from dbo.Users
) t
where t.num between ? and ?

