select id, name, client_type
from (
    select id, last_name + ', ' + first_name name, first_name + ' ' + last_name search_name, 'PERSON' client_type from dbo.Persons
    union all 
    select id, name, name, 'COMPANY' from dbo.Corporates
    union all
    select id, name, name, 'GROUP' from dbo.Groups
    union all
    select id, name, name, 'VILLAGE_BANK' from dbo.Villages
) t
where t.search_name like '%' + ? + '%'
order by name
