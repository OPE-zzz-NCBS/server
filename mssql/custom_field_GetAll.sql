select id,
    caption,
    [type],
    owner,
    tab,
    [unique],
    mandatory,
    [order],
    isnull(extra, '') extra
from dbo.CustomFields
where deleted = 0
order by tab, [order]
