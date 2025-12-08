from django.http import JsonResponse
from django.views.decorators.csrf import csrf_exempt
from django.conf import settings
import pymongo
import json
from datetime import datetime, timedelta
import jwt
from django.contrib.auth.hashers import make_password, check_password


client = pymongo.MongoClient(settings.DATABASES['default']['CLIENT']['host'])
db = client["SAMS"]
users_collection = db["Users"]


@csrf_exempt
def register(request):
    """
    Registers a new user.
    Hashes password and stores in MongoDB.
    Rejects duplicate emails.
    """
    if request.method != "POST":
        return JsonResponse({"error": "Invalid request method"}, status=405)

    try:
        data = json.loads(request.body)

        user_id = data.get("user_id")
        name = data.get("name")
        email = data.get("email")
        password = data.get("password")
        role = data.get("role", "student")

        # Validate required fields
        if not all([user_id, name, email, password]):
            return JsonResponse({"error": "Missing required fields"}, status=400)

        # Check if email already exists
        if users_collection.find_one({"email": email}):
            return JsonResponse({"error": "Email already exists"}, status=400)

        # Hash password
        hashed_password = make_password(password)

        new_user = {
            "user_id": user_id,
            "name": name,
            "email": email,
            "password": hashed_password,
            "role": role,
            "registered_at": datetime.now().isoformat()
        }

        users_collection.insert_one(new_user)

        return JsonResponse({"message": "User registered successfully"}, status=201)

    except Exception as e:
        return JsonResponse({"error": str(e)}, status=500)


@csrf_exempt
def login(request):
    """
    Validates user credentials.
    Returns a stateless JWT token with expiry.
    """
    if request.method != "POST":
        return JsonResponse({"error": "Invalid request method"}, status=405)

    try:
        data = json.loads(request.body)

        email = data.get("email")
        password = data.get("password")

        if not all([email, password]):
            return JsonResponse({"error": "Missing fields"}, status=400)

        # Get user from MongoDB
        user = users_collection.find_one({"email": email})
        if not user:
            return JsonResponse({"error": "Invalid credentials"}, status=401)

        # Validate password
        if not check_password(password, user["password"]):
            return JsonResponse({"error": "Invalid credentials"}, status=401)

        # BUILD JWT PAYLOAD
        expiry_minutes = getattr(settings, "JWT_EXP_MINUTES", 60)

        payload = {
            "user_id": user["user_id"],
            "email": user["email"],
            "role": user.get("role", "student"),
            "iat": datetime.utcnow(),
            "exp": datetime.utcnow() + timedelta(minutes=expiry_minutes)
        }

        # Sign token
        token = jwt.encode(payload, settings.SECRET_KEY, algorithm="HS256")

        return JsonResponse({
            "message": "Login successful",
            "token": token,
            "expires_in_seconds": expiry_minutes * 60,
            "user": {
                "user_id": user["user_id"],
                "name": user["name"],
                "email": user["email"],
                "role": user.get("role", "student")
            }
        }, status=200)

    except jwt.ExpiredSignatureError:
        return JsonResponse({"error": "Token expired"}, status=401)

    except jwt.InvalidTokenError:
        return JsonResponse({"error": "Invalid token"}, status=401)

    except Exception as e:
        return JsonResponse({"error": str(e)}, status=500)
