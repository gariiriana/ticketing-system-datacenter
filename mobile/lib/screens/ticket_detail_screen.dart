import 'package:flutter/material.dart';
import '../models/ticket.dart';
import '../services/api_service.dart';

class TicketDetailScreen extends StatefulWidget {
  final Ticket ticket;
  final bool isAdmin;

  const TicketDetailScreen({
    super.key,
    required this.ticket,
    required this.isAdmin,
  });

  @override
  State<TicketDetailScreen> createState() => _TicketDetailScreenState();
}

class _TicketDetailScreenState extends State<TicketDetailScreen> {
  final ApiService _apiService = ApiService();
  bool _isLoading = false;

  Future<void> _updateStatus(String newStatus) async {
    final confirm = await showDialog<bool>(
      context: context,
      builder: (ctx) => AlertDialog(
        backgroundColor: const Color(0xFF1E1E2E),
        title: Text(
          newStatus == 'approved' ? 'Approve Tiket?' : 'Reject Tiket?',
          style: const TextStyle(color: Colors.white),
        ),
        content: Text(
          'Apakah Anda yakin ingin ${newStatus == 'approved' ? 'menyetujui' : 'menolak'} tiket ini?',
          style: const TextStyle(color: Colors.white70),
        ),
        actions: [
          TextButton(
            onPressed: () => Navigator.pop(ctx, false),
            child: const Text('Batal'),
          ),
          ElevatedButton(
            onPressed: () => Navigator.pop(ctx, true),
            style: ElevatedButton.styleFrom(
              backgroundColor: newStatus == 'approved'
                  ? Colors.greenAccent
                  : Colors.redAccent,
            ),
            child: Text(newStatus == 'approved' ? 'Ya, Approve' : 'Ya, Reject',
                style: const TextStyle(color: Colors.black)),
          ),
        ],
      ),
    );

    if (confirm != true) return;

    setState(() => _isLoading = true);
    try {
      await _apiService.updateTicketStatus(widget.ticket.id, newStatus);
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text(
              newStatus == 'approved'
                  ? '✅ Tiket berhasil diapprove!'
                  : '❌ Tiket berhasil direject!',
            ),
            backgroundColor:
                newStatus == 'approved' ? Colors.green : Colors.redAccent,
          ),
        );
        Navigator.pop(context, true);
      }
    } catch (e) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
              content: Text('Gagal: $e'),
              backgroundColor: Colors.redAccent),
        );
      }
    } finally {
      if (mounted) setState(() => _isLoading = false);
    }
  }

  Color _statusColor(String status) {
    switch (status.toLowerCase()) {
      case 'pending':
        return Colors.orangeAccent;
      case 'approved':
        return Colors.greenAccent;
      case 'rejected':
        return Colors.redAccent;
      case 'in_progress':
        return Colors.blueAccent;
      default:
        return Colors.grey;
    }
  }

  @override
  Widget build(BuildContext context) {
    final ticket = widget.ticket;
    final statusColor = _statusColor(ticket.status);
    final canAct = widget.isAdmin && ticket.status == 'pending';

    return Scaffold(
      appBar: AppBar(
        title: const Text('Detail Tiket'),
        leading: IconButton(
          icon: const Icon(Icons.arrow_back_ios_rounded),
          onPressed: () => Navigator.pop(context),
        ),
      ),
      body: SingleChildScrollView(
        padding: const EdgeInsets.all(20),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            // Status Badge
            Container(
              width: double.infinity,
              padding: const EdgeInsets.all(20),
              decoration: BoxDecoration(
                gradient: LinearGradient(
                  colors: [
                    statusColor.withOpacity(0.2),
                    statusColor.withOpacity(0.05),
                  ],
                ),
                borderRadius: BorderRadius.circular(16),
                border: Border.all(color: statusColor.withOpacity(0.4)),
              ),
              child: Column(
                children: [
                  CircleAvatar(
                    radius: 30,
                    backgroundColor: statusColor.withOpacity(0.2),
                    child: Icon(
                      ticket.status == 'approved'
                          ? Icons.check_circle_rounded
                          : ticket.status == 'rejected'
                              ? Icons.cancel_rounded
                              : ticket.status == 'in_progress'
                                  ? Icons.engineering_rounded
                                  : Icons.hourglass_empty_rounded,
                      color: statusColor,
                      size: 32,
                    ),
                  ),
                  const SizedBox(height: 12),
                  Text(
                    ticket.status.toUpperCase(),
                    style: TextStyle(
                      color: statusColor,
                      fontSize: 18,
                      fontWeight: FontWeight.bold,
                      letterSpacing: 2,
                    ),
                  ),
                ],
              ),
            ),
            const SizedBox(height: 24),

            // Info Cards
            _infoCard(
              icon: Icons.description_rounded,
              label: 'Deskripsi Kendala',
              value: ticket.description,
            ),
            const SizedBox(height: 12),
            _infoCard(
              icon: Icons.location_on_rounded,
              label: 'Site',
              value: ticket.siteId,
            ),
            const SizedBox(height: 12),
            _infoCard(
              icon: Icons.access_time_rounded,
              label: 'Dibuat Pada',
              value: '${ticket.createdAt.day}/${ticket.createdAt.month}/${ticket.createdAt.year} '
                  '${ticket.createdAt.hour}:${ticket.createdAt.minute.toString().padLeft(2, '0')}',
            ),
            const SizedBox(height: 12),
            _infoCard(
              icon: Icons.tag_rounded,
              label: 'Ticket ID',
              value: ticket.id,
            ),

            // Admin action buttons
            if (canAct) ...[
              const SizedBox(height: 32),
              const Text(
                'Tindakan Admin',
                style: TextStyle(
                  color: Colors.white54,
                  fontWeight: FontWeight.bold,
                  fontSize: 14,
                  letterSpacing: 1,
                ),
              ),
              const SizedBox(height: 16),
              Row(
                children: [
                  Expanded(
                    child: _isLoading
                        ? const Center(
                            child: CircularProgressIndicator(
                                color: Color(0xFF6C63FF)))
                        : OutlinedButton.icon(
                            onPressed: () => _updateStatus('rejected'),
                            icon: const Icon(Icons.close_rounded,
                                color: Colors.redAccent),
                            label: const Text('Tolak',
                                style: TextStyle(color: Colors.redAccent)),
                            style: OutlinedButton.styleFrom(
                              side:
                                  const BorderSide(color: Colors.redAccent),
                              padding:
                                  const EdgeInsets.symmetric(vertical: 14),
                              shape: RoundedRectangleBorder(
                                  borderRadius: BorderRadius.circular(12)),
                            ),
                          ),
                  ),
                  const SizedBox(width: 12),
                  Expanded(
                    child: _isLoading
                        ? const SizedBox()
                        : ElevatedButton.icon(
                            onPressed: () => _updateStatus('approved'),
                            icon: const Icon(Icons.check_rounded),
                            label: const Text('Setujui'),
                            style: ElevatedButton.styleFrom(
                              backgroundColor: Colors.greenAccent,
                              foregroundColor: Colors.black,
                              padding:
                                  const EdgeInsets.symmetric(vertical: 14),
                              shape: RoundedRectangleBorder(
                                  borderRadius: BorderRadius.circular(12)),
                            ),
                          ),
                  ),
                ],
              ),
            ],

            if (ticket.status != 'pending' && widget.isAdmin) ...[
              const SizedBox(height: 24),
              Container(
                width: double.infinity,
                padding: const EdgeInsets.all(16),
                decoration: BoxDecoration(
                  color: const Color(0xFF1E1E2E),
                  borderRadius: BorderRadius.circular(12),
                ),
                child: const Text(
                  'Tiket ini sudah diproses dan tidak bisa diubah lagi.',
                  textAlign: TextAlign.center,
                  style: TextStyle(color: Colors.white38, fontSize: 13),
                ),
              ),
            ],
          ],
        ),
      ),
    );
  }

  Widget _infoCard(
      {required IconData icon, required String label, required String value}) {
    return Container(
      padding: const EdgeInsets.all(16),
      decoration: BoxDecoration(
        color: const Color(0xFF1E1E2E),
        borderRadius: BorderRadius.circular(12),
      ),
      child: Row(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Container(
            padding: const EdgeInsets.all(8),
            decoration: BoxDecoration(
              color: const Color(0xFF6C63FF).withOpacity(0.15),
              borderRadius: BorderRadius.circular(8),
            ),
            child: Icon(icon, color: const Color(0xFF6C63FF), size: 20),
          ),
          const SizedBox(width: 12),
          Expanded(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(label,
                    style:
                        const TextStyle(color: Colors.white38, fontSize: 12)),
                const SizedBox(height: 4),
                Text(value,
                    style: const TextStyle(
                        color: Colors.white, fontSize: 15)),
              ],
            ),
          ),
        ],
      ),
    );
  }
}
