const std = @import("std");

pub const SignUp = extern struct {
    Username: [*]const u8,
    Bio: [*]const u8,
    Surname: [*]const u8,
    Firstname: [*]const u8,
    Lastname: [*]const u8,
    Email: [*]const u8,
    Password: [*]const u8,

    fn hash(self: @This()) []u8 {
        const hashstring = try std.crypto.pwhash.bcrypt.strHash(self.Password);

        return hashstring;
    }

    pub export fn string(self: @This()) []u8 {
        const allocator = std.heap.GeneralPurposeAllocator;
        return try std.fmt.allocPrint(allocator, "{{\"username\":\"{s}\",\"bio\":{s},\"surname\":{s},\"firstname\":{s},\"lastname\":{s},\"email\":{s}, ,\"password\":{s}}}", .{ self.Username, self.Bio, self.Surname, self.Firstname, self.Lastname, self.Email, self.hash() });
    }
};

pub const LogIn = extern struct {
    Email: [*]const u8,
    Password: [*]const u8,

    pub export fn string(self: @This()) []u8 {
        const allocator = std.heap.GeneralPurposeAllocator;
        return try std.fmt.allocPrint(allocator, "{{\"email\":{s}, \"password\":{s}}}", .{ self.Email, self.Password });
    }
};

pub const LogOut = extern struct {
    Id: i64,

    pub export fn string(self: @This()) []u8 {
        const allocator = std.heap.GeneralPurposeAllocator;
        return try std.fmt.allocPrint(allocator, "{{\"id\":{d}}}", .{self.Id});
    }
};

pub const ChangeEmail = extern struct {
    Id: i64,
    Email: [*]const u8,

    pub export fn string(self: @This()) []u8 {
        const allocator = std.heap.GeneralPurposeAllocator;
        return try std.fmt.allocPrint(allocator, "{{\"id\":{d}, \"email\":{s}}}", .{ self.Id, self.Email });
    }
};

pub const ResetPassword = extern struct {
    Id: i64,
    Password: [*]const u8,

    pub export fn string(self: @This()) []u8 {
        const allocator = std.heap.GeneralPurposeAllocator;
        return try std.fmt.allocPrint(allocator, "{{\"id\":{d}, \"password\":{s}}}", .{ self.Id, self.Password });
    }
};

pub const UpdateProfile = extern struct {
    Username: [*]const u8,
    Bio: [*]const u8,

    pub export fn string(self: @This()) []u8 {
        const allocator = std.heap.GeneralPurposeAllocator;
        return try std.fmt.allocPrint(allocator, "{{\"username\":{s}, \"bio\":{s}}}", .{ self.Username, self.Bio });
    }
};

pub const CreateCategory = extern struct {
    UserId: i64,
    Name: [*]const u8,

    pub export fn string(self: @This()) []u8 {
        const allocator = std.heap.GeneralPurposeAllocator;
        return try std.fmt.allocPrint(allocator, "{{\"userid\":{d}, \"name\":{s}}}", .{ self.UserId, self.Name });
    }
};

pub const UpdateCategory = extern struct {
    Name: [*]const u8,

    pub export fn string(self: @This()) []u8 {
        const allocator = std.heap.GeneralPurposeAllocator;
        return try std.fmt.allocPrint(allocator, "{{\"name\":{s}}}", .{self.Name});
    }
};

pub const CreateBlog = extern struct {
    UserId: i64,
    Title: [*]const u8,
    Body: [*]const u8,
    Categories: [*]const u8,

    pub export fn string(self: @This()) []u8 {
        const allocator = std.heap.GeneralPurposeAllocator;
        return try std.fmt.allocPrint(allocator, "{{\"userid\":{d}, \"title\":{s}, \"bodies\":{s}, \"categories\":{s}}}", .{ self.UserId, self.Title, self.Body, self.Categories });
    }
};

pub const UpdateBlog = extern struct {
    Title: [*]const u8,
    Body: [*]const u8,

    pub export fn string(self: @This()) []u8 {
        const allocator = std.heap.GeneralPurposeAllocator;
        return try std.fmt.allocPrint(allocator, "{{\"title\":{s}, \"bodies\":{s}}}", .{ self.Title, self.Body });
    }
};
