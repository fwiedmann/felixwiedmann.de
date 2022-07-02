package site


allow {
    input.method = "GET"
    input.path = ["opinions", opinion_id]
    security_filter
}

security_filter{
    data.opinions.ownerId == input.ownerId
}