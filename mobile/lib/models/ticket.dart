class Ticket {
  final String id;
  final String description;
  final String status;
  final String siteId;
  final String photoUrl;
  final DateTime createdAt;

  Ticket({
    required this.id,
    required this.description,
    required this.status,
    required this.siteId,
    required this.photoUrl,
    required this.createdAt,
  });

  factory Ticket.fromJson(Map<String, dynamic> json) {
    return Ticket(
      id: json['id'],
      description: json['description'],
      status: json['status'],
      siteId: json['site_id'],
      photoUrl: json['photo_url'] ?? '',
      createdAt: DateTime.parse(json['created_at']),
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'description': description,
      'status': status,
      'site_id': siteId,
      'photo_url': photoUrl,
      'created_at': createdAt.toIso8601String(),
    };
  }
}
