mutation fetchUser($id: ID!) {
    getUser(id: $id) { // ここの名前は上と同じでなくてもいい
      id
    }
  }