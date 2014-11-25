select c.id,
    case
        when num.num = 1 then c.name
        else c.name + ' (' + d.name + ')'
    end name,
    d.id
from dbo.City c
left join
(
    select name, count(id) num
    from dbo.City
    where deleted = 0
    group by name
) num on num.name = c.name
left join dbo.Districts d on d.id = c.district_id
where c.deleted = 0
order by c.name
