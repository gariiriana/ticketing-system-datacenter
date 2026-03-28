import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:firebase_auth/firebase_auth.dart';
import '../models/ticket.dart';

class ApiService {
  final String baseUrl = "http://localhost:8080/api/v1"; // Change to actual server IP for real device

  Future<String?> _getToken() async {
    return await FirebaseAuth.instance.currentUser?.getIdToken();
  }

  Future<List<Ticket>> getTickets() async {
    final token = await _getToken();
    final response = await http.get(
      Uri.parse("$baseUrl/tickets"),
      headers: {
        "Authorization": "Bearer $token",
      },
    );

    if (response.statusCode == 200) {
      List jsonResponse = json.decode(response.body);
      return jsonResponse.map((data) => Ticket.fromJson(data)).toList();
    } else {
      throw Exception('Failed to load tickets');
    }
  }

  Future<void> createTicket(Ticket ticket) async {
    final token = await _getToken();
    final response = await http.post(
      Uri.parse("$baseUrl/tickets"),
      headers: {
        "Content-Type": "application/json",
        "Authorization": "Bearer $token",
      },
      body: jsonEncode(ticket.toJson()),
    );

    if (response.statusCode != 201) {
      throw Exception('Failed to create ticket');
    }
  }
}
