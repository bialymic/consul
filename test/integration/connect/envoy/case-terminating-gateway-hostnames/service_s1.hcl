services {
  name = "s1"
  port = 8080
  connect {
    sidecar_service {
      proxy {
        upstreams = [
          {
            destination_name = "s4"
            local_bind_port = 5000
          }
        ]
      }
    }
  }
}