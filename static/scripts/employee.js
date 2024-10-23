document.addEventListener('DOMContentLoaded', function() {
    var calendarEl = document.getElementById('calendar');
    
    var calendar = new FullCalendar.Calendar(calendarEl, {
        initialView: 'dayGridMonth',
        events: '/api/employee/leaves', 
        headerToolbar: {
            left: 'prev,next today',
            center: 'title',
            right: 'dayGridMonth,timeGridWeek,timeGridDay'
        }
    });

    calendar.render();

    document.getElementById('update-employee-form').addEventListener('submit', function(event) {
        event.preventDefault(); 
        const formData = new FormData(this);
        fetch('/updateEmployee', { 
            method: 'POST',
            body: formData
        }).then(response => {
            if (response.ok) {
                window.location.href = '/employee?uuid=' + formData.get('employeeUuid');
            } else {
                alert('Erreur lors de la mise à jour de l\'employé: ' + response.statusText);
            }
        }).catch(error => {
            alert('Erreur lors de la mise à jour de l\'employé: ' + error.message);
        });
    });

    document.getElementById('delete-employee').addEventListener('click', function() {
        const employeeUuid = prompt('Enter employee UUID to delete:');
        if (employeeUuid) {
            const formData = new FormData();
            formData.append('employeeUuid', employeeUuid);
            formData.append('confirm', 'true');

            fetch('/deleteEmployee', {
                method: 'POST',
                body: formData
            }).then(response => {
                if (response.ok) {
                    window.location.href = '/home'; 
                } else {
                    alert('Erreur lors de la suppression de l\'employé: ' + response.statusText);
                }
            });
        }
    });
});