module basic-chapter

go 1.15

require golang.org/x/tour v0.0.0-20201207214521-004403599411

replace local_module  => ./local_module
require local_module v0.0.0-00010101000000-000000000000

