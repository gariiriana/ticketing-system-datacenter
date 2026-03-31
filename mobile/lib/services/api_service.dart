import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:firebase_auth/firebase_auth.dart';
import '../models/ticket.dart';

class ApiService {
  /// For Android physical device: use your PC's local IP (e.g. 192.168.1.x)
  /// For Android emulator: use 10.0.2.2
  /// For web/desktop: use localhost
  static const String _baseUrl = 'http://10.10.20.153:8080/api/v1';

  Future<String?> _getToken() async {
    return await FirebaseAuth.instance.currentUser?.getIdToken();
  }

  Future<void> syncProfile() async {
    final token = await _getToken();
    if (token == null) throw Exception('Tidak terautentikasi');

    final response = await http
        .get(
          Uri.parse('$_baseUrl/user/sync'),
          headers: _authHeaders(token),
        )
        .timeout(const Duration(seconds: 15));

    if (response.statusCode != 200) {
      throw Exception('Gagal sinkronisasi profile: ${response.body}');
    }
  }

  Map<String, String> _authHeaders(String token) => {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer $token',
      };

  Future<List<Ticket>> getTickets() async {
    final token = await _getToken();
    if (token == null) throw Exception('Tidak terautentikasi');

    final response = await http
        .get(
          Uri.parse('$_baseUrl/tickets'),
          headers: _authHeaders(token),
        )
        .timeout(const Duration(seconds: 15));

    if (response.statusCode == 200) {
      final List json = jsonDecode(response.body);
      return json.map((d) => Ticket.fromJson(d)).toList();
    } else {
      throw Exception(
          'Gagal memuat tiket (${response.statusCode}): ${response.body}');
    }
  }

  Future<void> createTicket(Ticket ticket) async {
    final token = await _getToken();
    if (token == null) throw Exception('Tidak terautentikasi');

    final response = await http
        .post(
          Uri.parse('$_baseUrl/tickets'),
          headers: _authHeaders(token),
          body: jsonEncode({
            'description': ticket.description,
            'site_id': ticket.siteId,
          }),
        )
        .timeout(const Duration(seconds: 15));

    if (response.statusCode != 201) {
      throw Exception(
          'Gagal membuat tiket (${response.statusCode}): ${response.body}');
    }
  }

  Future<void> updateTicketStatus(String ticketId, String status,
      {String? reason}) async {
    final token = await _getToken();
    if (token == null) throw Exception('Tidak terautentikasi');

    final response = await http
        .patch(
          Uri.parse('$_baseUrl/tickets/$ticketId'),
          headers: _authHeaders(token),
          body: jsonEncode({
            'status': status,
            'reason': reason ?? '',
          }),
        )
        .timeout(const Duration(seconds: 15));

    if (response.statusCode != 200) {
      throw Exception(
          'Gagal update status (${response.statusCode}): ${response.body}');
    }
  }
}
