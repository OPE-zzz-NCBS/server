select cf.id,
    cf.caption,
    cf.[type],
    cf.owner,
    cf.tab,
    cf.[unique],
    cf.mandatory,
    cf.[order],
    isnull(cf.extra, '') extra,
    isnull(cfv.value, '') value
from dbo.CustomFields cf
left join
(
    select value, field_id
    from dbo.CustomFieldsValues
    where owner_id = ?
) cfv on cfv.field_id = cf.id
where cf.owner = 'Person' and cf.deleted = 0
order by cf.[order]
